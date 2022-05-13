package controllers

import (
	"github.com/astaxie/beego"
	"movie/backend/models"
	"time"
)

type CommentController struct {
	beego.Controller
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      time.Time       `json:"addTime"`
	AddTimeTitle string          `json:"AddTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
	EpisodesId   int 			 `json:"episodesId"`
}


//获取评论列表
// @router /comment/list [*]
func (c *CommentController) List(){
	episodesId,_ := c.GetInt("episodesId")   //剧集id
	limit,_ := c.GetInt("limit")
	offset,_ := c.GetInt("offset")

	if episodesId == 0 {
		c.Data["json"] = ReturnError(4001, "必须指定视频剧集")
		c.ServeJSON()
		return
	}
	if limit == 0 {
		limit = 12
	}
	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err == nil {
		var data []CommentInfo
		var commentInfo CommentInfo
		//
		////获取uid channel
		//uidChan := make(chan int, 12)
		//closeChan := make(chan bool, 5)
		//resChan := make(chan models.UserInfo, 12)
		////把获取到的uid放到channel中
		//go func() {
		//	for _, v := range comments {
		//		uidChan <- v.UserId
		//	}
		//	close(uidChan)
		//}()
		////处理uidChannel中的信息
		//for i := 0; i < 5; i++ {
		//	go chanGetUserInfo(uidChan, resChan, closeChan)
		//}
		////判断是否执行完成，信息聚合
		//go func() {
		//	for i := 0; i < 5; i++ {
		//		<-closeChan
		//	}
		//	close(resChan)
		//	close(closeChan)
		//}()

		//userInfoMap := make(map[int]models.UserInfo)
		//for r := range resChan {
		//	userInfoMap[r.Id] = r
		//}
		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DataFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount
			commentInfo.EpisodesId = v.EpisodesId
			//获取用户信息
			commentInfo.UserInfo,_ = models.GetUserInfo(v.UserId)
			data = append(data, commentInfo)
		}

		c.Data["json"] = ReturnSuccess(0, "success", data, num)
		c.ServeJSON()
	} else {
		c.Data["json"] = ReturnError(4004, "没有相关内容")
		c.ServeJSON()
		return
	}
}

//func chanGetUserInfo(uidChan chan int, resChan chan models.UserInfo, closeChan chan bool) {
//	for uid := range uidChan {
//		res, err := models.GetUserInfo(uid)
//		fmt.Println(res)
//		if err == nil {
//			resChan <- res
//		}
//	}
//	closeChan <- true
//}

// 保存评论
// @router /comment/save [*]
func (c *CommentController) Save(){
	content := c.GetString("content")
	uid,_ := c.GetInt("uid")
	episodesId,_ := c.GetInt("episodesId")
	videoId,_ := c.GetInt("videoId")
	if content == "" {
		c.Data["json"] = ReturnError(4001, "内容不能为空")
		c.ServeJSON()
		return
	}
	if uid == 0 {
		c.Data["json"] = ReturnError(4002, "请先登录")
		c.ServeJSON()
		return
	}
	if episodesId == 0 {
		c.Data["json"] = ReturnError(4003, "必须指定频道剧集id")
		c.ServeJSON()
		return
	}
	if videoId == 0 {
		c.Data["json"] = ReturnError(4004, "必须指定视频id")
		c.ServeJSON()
		return
	}
	err := models.SaveComment(content, uid, episodesId, videoId)
	if err == nil {
		c.Data["json"] = ReturnSuccess(0, "success", "", 1)
		c.ServeJSON()
	}else {
		c.Data["json"] = ReturnError(5000, err)
		c.ServeJSON()
		return
	}
}
