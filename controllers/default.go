package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/tfxingApiProxy/models"
	)

type Request struct {
	Optid    string
	Version  interface{}
	Terminal interface{}
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {

	req := Request{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if nil != err {
		log.Println("请求参数无法解析", err.Error())
	} else {
		file, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(file)
		index := strings.LastIndex(path, string(os.PathSeparator))
		rootPath := path[:index] + "/"
		parentFolder := "jsons"
		url := c.Ctx.Request.URL.Path
		fileName := req.Optid + ".json"
		filePath := filepath.Join(rootPath, parentFolder, url, fileName)
		log.Println("文件路径" + filePath)

		data, err := ioutil.ReadFile(filePath)
		if nil != err {
			log.Println("读取回送信息文件错误", err.Error())
			log.Println("准备重定向到真实服务器")
			//直接转发请求
			c.Redirect(models.Config.RealServer+c.Ctx.Request.URL.RequestURI(), 302)
		} else {
			var jsonObj map[string]interface{}
			err := json.Unmarshal(data, &jsonObj)
			if nil != err {
				log.Println("解析回送信息文件错误，也许是json格式错误.", err.Error())
			} else {
				c.Data["json"] = jsonObj
				c.ServeJSON()
			}
		}
	}
}
