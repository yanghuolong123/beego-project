package controllers

import (
	"bufio"
	"go-web/components/ext"
	"go-web/components/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
	prefix := "static/uploads/"
	part := prefix + filename + "_" + chunk + ".part"
	this.SaveToFile("file", part)
	count, err := strconv.Atoi(chunks)
	redis := utils.Redis()
	redis.Incr(filename)
	num, err := strconv.Atoi(string(redis.Get(filename).([]uint8)))
	if num == count {
		redis.Delete(filename)
		log.Println("==================== num:", num)
		outfile := prefix + filename
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
	this.SendResJsonp(0, "ok", filename+":::"+chunks)
}
