package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"weixin/goServer/models"
)

type LoginController struct {
	beego.Controller
}

func(this *LoginController) Post(){
	log.Print("login-------post");
	beego.Notice("login-------post");
}


type Student struct {
	Name string
	Age int
	Code string
	Url string
}
func(this *LoginController) Get(){
	//log.Print("login-------Get");
	//beego.Notice("login-------Get");
	//this.Ctx.WriteString("hello")


	//TODO 获取code
	code:=this.GetString("code");
	name:=this.GetString("name");
	age,_:=this.GetInt("age");
	if code != "" {
		mystruct:=Student{
			Name:name,
			Age:age,
			Code:code,
			Url:"./static/img/tuoyuan.png",
		}
		this.Data["json"] = &mystruct
		this.ServeJSON()

		return
	}

	//TODO 向微信后台请求相关数据


	//TODO


	//http请求url
	resp,err:=http.Get("http://localhost:8080/static/img/tuoyuan.png");
	if err != nil {
		beego.Error("login-------http-----get");
		this.Ctx.WriteString("")
		return;
	}
	defer resp.Body.Close();



	//文件夹不存在，创建文件夹
	if !models.Exists ("./player"){
		os.Mkdir("./player",os.ModePerm  )
	}

	//请求成功,写入图片
	if resp.StatusCode == http.StatusOK {
		body,_:=ioutil.ReadAll(resp.Body);
		id:=1001;
		out,_:=os.Create(fmt.Sprintf("./player/%d.png",id));
		io.Copy(out,bytes.NewReader(body))
	}
	this.Ctx.WriteString(fmt.Sprintf("StatusCode:%v",resp.StatusCode));
}





