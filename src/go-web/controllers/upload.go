package controllers

import (
	"bufio"
	"fmt"
	"go-web/components/ext"
	"go-web/components/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type UploadController struct {
	ext.BaseController
}

func (this *UploadController) Webupload() {
	method := this.Ctx.Input.Method()
	utils.Logs().Info("method:" + method)
	if method == "OPTIONS" {
		this.StopRun()
	}
	filename := this.Ctx.Input.Query("name")
	chunks := this.Ctx.Input.Query("chunks")
	chunk := this.Ctx.Input.Query("chunk")
	f, h, err := this.GetFile("file")
	if err != nil {
		utils.Logs().Info(err.Error())
	}
	defer f.Close()
	ext := filepath.Ext(h.Filename)
	filename = utils.Md5(filename) + ext
	prefix := "static/uploads/tmp/"
	part := prefix + filename + "_" + chunk + ".part"
	this.SaveToFile("file", part)
	count, err := strconv.Atoi(chunks)
	redis := utils.Redis()
	redis.Incr(filename)
	num, err := strconv.Atoi(string(redis.Get(filename).([]uint8)))
	y, m, d := utils.Date()
	dir := "uploads/" + fmt.Sprintf("%d/%d/%d/", y, m, d)
	if num == count {
		redis.Delete(filename)
		log.Println("==================== num:", num)
		if !utils.PathExist(dir) {
			os.MkdirAll(dir, os.ModePerm)
		}
		outfile := "static/" + dir + fmt.Sprintf("%d%d", time.Now().Unix(), utils.RandNum(10000, 99999)) + ext
		out, _ := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY, 0600)
		bWriter := bufio.NewWriter(out)
		for i := 0; i < count; i++ {
			infile := prefix + filename + "_" + strconv.Itoa(i) + ".part"
			in, _ := os.Open(infile)
			bReader := bufio.NewReader(in)
			for {
				buffer := make([]byte, 1024)
				readCount, err := bReader.Read(buffer)
				if err == io.EOF {
					os.Remove(infile)
					break
				} else {
					bWriter.Write(buffer[:readCount])
				}
			}
			log.Println("==================== i:", i)
		}
		bWriter.Flush()
	}

	utils.Logs().Info("filename:" + filename + " chunks:" + chunks + " chunk:" + chunk)
	this.SendResJsonp(0, "ok", dir+filename)
}
