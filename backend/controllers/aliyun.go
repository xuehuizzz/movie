package controllers
//
//import (
//	"encoding/json"
//	"fmt"
//	"movie/backend/models"
//	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
//	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
//	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
//
//	"github.com/astaxie/beego"
//)
//
//type AliyunController struct {
//	beego.Controller
//}
//
//var (
//	accessKeyId     = "LTAI4Fy6F7uJvhN2iGGxqtXH"
//	accessKeySecret = "vifM2gxcEiLj1ZPxjlR2RZI9ln4VB2"
//)
//
//type JSONS struct {
//	RequestId     string
//	UploadAddress string
//	UploadAuth    string
//	VideoId       string
//}
//
//func (a *AliyunController) InitVodClient(accessKeyId string, accessKeySecret string) (client *vod.Client, err error) {
//	// 点播服务接入区域
//	regionId := "cn-shanghai"
//	// 创建授权对象
//	credential := &credentials.AccessKeyCredential{
//		accessKeyId,
//		accessKeySecret,
//	}
//	// 自定义config
//	config := sdk.NewConfig()
//	config.AutoRetry = true     // 失败是否自动重试
//	config.MaxRetryTime = 3     // 最大重试次数
//	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
//	// 创建vodClient实例
//	return vod.NewClientWithOptions(regionId, config, credential)
//}
//
//func (a *AliyunController) MyCreateUploadVideo(client *vod.Client, title string, desc string, fileName string, coverUrl string, tags string) (response *vod.CreateUploadVideoResponse, err error) {
//	request := vod.CreateCreateUploadVideoRequest()
//	request.Title = title
//	request.Description = desc
//	request.FileName = fileName
//	//request.CateId = "-1"
//	request.CoverURL = coverUrl
//	request.Tags = tags
//	request.AcceptFormat = "JSON"
//	return client.CreateUploadVideo(request)
//}
//
//// @router /aliyun/create/upload/video [*]
//func (a *AliyunController) CreateUploadVideo() {
//	title := a.GetString("title")
//	desc := a.GetString("desc")
//	fileName := a.GetString("fileName")
//	coverUrl := a.GetString("coverUrl")
//	tags := a.GetString("tags")
//	client, err := a.InitVodClient(accessKeyId, accessKeySecret)
//	if err != nil {
//		panic(err)
//	}
//	response, err := a.MyCreateUploadVideo(client, title, desc, fileName, coverUrl, tags)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(response.GetHttpContentString())
//	//fmt.Printf("VideoId: %s\n UploadAddress: %s\n UploadAuth: %s",
//	//    response.VideoId, response.UploadAddress, response.UploadAuth)
//	data := &JSONS{
//		response.RequestId,
//		response.UploadAddress,
//		response.UploadAuth,
//		response.VideoId,
//	}
//	a.Data["json"] = data
//	a.ServeJSON()
//}
//
//func (a *AliyunController) MyRefreshUploadVideo(client *vod.Client, videoId string) (response *vod.RefreshUploadVideoResponse, err error) {
//	request := vod.CreateRefreshUploadVideoRequest()
//	request.VideoId = videoId
//	request.AcceptFormat = "JSON"
//	return client.RefreshUploadVideo(request)
//}
//
//// @router /aliyun/refresh/upload/video [*]
//func (a *AliyunController) RefreshUploadVideo() {
//	videoId := a.GetString("videoId")
//
//	client, err := a.InitVodClient(accessKeyId, accessKeySecret)
//	if err != nil {
//		panic(err)
//	}
//	response, err := a.MyRefreshUploadVideo(client, videoId)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(response.GetHttpContentString())
//	//fmt.Printf("UploadAddress: %s\n UploadAuth: %s", response.UploadAddress, response.UploadAuth)
//	data := &JSONS{
//		response.RequestId,
//		response.UploadAddress,
//		response.UploadAuth,
//		response.VideoId,
//	}
//	a.Data["json"] = data
//	a.ServeJSON()
//}
//
//func (a *AliyunController) MyGetPlayAuth(client *vod.Client, videoId string) (response *vod.GetVideoPlayAuthResponse, err error) {
//	request := vod.CreateGetVideoPlayAuthRequest()
//	request.VideoId = videoId
//	request.AcceptFormat = "JSON"
//	return client.GetVideoPlayAuth(request)
//}
//
//type PlayJSONS struct {
//	PlayAuth string
//}
//
//// @router /aliyun/video/play/auth [*]
//func (a *AliyunController) GetPlayAuth() {
//	videoId := a.GetString("videoId")
//	client, err := a.InitVodClient(accessKeyId, accessKeySecret)
//	if err != nil {
//		panic(err)
//	}
//	response, err := a.MyGetPlayAuth(client, videoId)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(response.GetHttpContentString())
//	//fmt.Printf("%s: %s\n", response.VideoMeta, response.PlayAuth)
//	data := &PlayJSONS{
//		response.PlayAuth,
//	}
//	a.Data["json"] = data
//	a.ServeJSON()
//}
//
//type CallbackData struct {
//	EventTime   string
//	EventType   string
//	VideoId     string
//	Status      string
//	Exteng      string
//	StreamInfos []CallbackStreamInfosData
//}
//type CallbackStreamInfosData struct {
//	Status     string
//	Bitrate    int
//	Definition string
//	Duration   int
//	Encrypt    bool
//	FileUrl    string
//	Format     string
//	Fps        int
//	Height     int
//	Size       int
//	Width      int
//	JobId      string
//}
//
////回调函数
//// @router /aliyun/video/callback [*]
//func (a *AliyunController) VideoCallback() {
//	var ob CallbackData
//	r := a.Ctx.Input.RequestBody
//	json.Unmarshal(r, &ob)
//
//	models.SaveAliyunVideo(ob.VideoId, string(r))
//	a.Ctx.WriteString("success")
//}
