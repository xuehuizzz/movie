package controllers

import (
	"crypto/md5"
	"fmt"
	"time"

	//"encoding/hex"
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code int `json:"code"`
	Msg interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64 `json:"count"`
}

func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) (json *JsonStruct){
	json = &JsonStruct{Code:code, Msg:msg, Items:items, Count:count}
	return
}

func ReturnError(code int, msg interface{}) (json *JsonStruct){
	json = &JsonStruct{Code:code, Msg:msg}
	return
}

//用户密码加密
func MD5V(password string) string {
	str_to_byte := []byte(password)
	byte_res := md5.Sum(str_to_byte)
	res := fmt.Sprintf("%x", byte_res)
	return res
}


//格式化时间
func DataFormat(t time.Time) string {
	video_time := t.Format("2006-01-02")
	return video_time
}

//字符串转日期
func str2time(s string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return t
}