package controllers

import (
	"go-web/components/ext"
	"go-web/components/utils"
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
	utils.Logs().Info("file:" + h.Filename)
	this.SaveToFile("file", "static/uploads/"+h.Filename+"_"+chunk+".part")

	utils.Logs().Info("filename:" + filename + " chunks:" + chunks + " chunk:" + chunk)
	this.SendResJsonp(0, "ok", filename+":::"+chunks)
}
