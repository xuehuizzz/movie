package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"movie/backend/models"
)

type VideoController struct {
	beego.Controller
}

//频道页-获取顶部广告
// @router /channel/advert [*]
func (v *VideoController) ChannelAdvert() {
	channelId, _ := v.GetInt("channelId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelAdvert(channelId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "数据请求失败, 请稍后再试~")
		v.ServeJSON()
		return
	}
}

//频道页-获取正在热播
// @router /channel/hot [*]
func (v *VideoController) ChannelHotList() {
	channelId, _ := v.GetInt("channelId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelHotList(channelId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
		return
	}
}

//频道页-根据频道地区获取推荐的视频
// @router /channel/recommend/region [*]
func (v *VideoController) ChannelRecommendRegionlist() {
	channelId, _ := v.GetInt("channelId")
	regionId, _ := v.GetInt("regionId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
		return
	}
	if regionId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道地区")
		v.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
		return
	}
}

//频道页-获取少女推荐
// @router /channel/recommend/type [*]
func (v *VideoController) ChannelRecommendTypeList() {
	channelId, _ := v.GetInt("channelId")
	typeId, _ := v.GetInt("typeId")
	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
		return
	}
	if typeId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道类型")
		v.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
		return
	}
}

//根据传入参数获取视频列表
// @router /channel/video [*]
func (v *VideoController) ChannelVideo() {
	channelId, _ := v.GetInt("channelId") //获取频道id
	regionId, _ := v.GetInt("regionId")   //获取频道地区id
	typeId, _ := v.GetInt("typeId")       //获取频道类型id
	end := v.GetString("end")             //状态
	sort := v.GetString("sort")           //排序
	limit, _ := v.GetInt("limit")         //页码
	offset, _ := v.GetInt("offset")

	if channelId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定频道")
		v.ServeJSON()
		return
	}
	if limit == 0 {
		limit = 12
	}
	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
		return
	}
}

//我的视频管理
// @router /user/video [*]
func (v *VideoController) UserVideo() {
	uid, _ := v.GetInt("uid")
	if uid == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定用户")
		v.ServeJSON()
		return
	}
	fmt.Println(uid)
	num, videos, err := models.GetUserVideo(uid)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", videos, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "没有相关内容")
		v.ServeJSON()
		return
	}
}

//获取视频详情
// @router /video/info [*]
func (v *VideoController) VideoInfo() {
	videoId, _ := v.GetInt("videoId")
	if videoId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定视频ID")
		v.ServeJSON()
		return
	}
	video, err := models.GetVideoInfo(videoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", video, 1)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败, 请稍后再试~")
		v.ServeJSON()
		return
	}
}

//获取剧集列表
// @router /video/episodes/list [*]
func (v *VideoController) VideoEpisodesList() {
	videoId, _ := v.GetInt("videoId")
	if videoId == 0 {
		v.Data["json"] = ReturnError(4001, "必须指定视频ID")
		v.ServeJSON()
		return
	}
	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(4004, "请求数据失败, 请稍后再试~")
		v.ServeJSON()
		return
	}
}

//保存用户上传视频信息
// @router /video/save [*]
func (v *VideoController) VideoSave() {
	playUrl := v.GetString("playUrl")
	title := v.GetString("title")
	subTitle := v.GetString("subTitle")
	channelId, _ := v.GetInt("channelId")
	typeId, _ := v.GetInt("typeId")
	regionId, _ := v.GetInt("regionId")
	uid, _ := v.GetInt("uid")
	aliyunVideoId := v.GetString("aliyunVideoId")

	if uid == 0 {
		v.Data["json"] = ReturnError(4001, "请先登录再继续~")
		v.ServeJSON()
		return
	}
	if playUrl == "" {
		v.Data["json"] = ReturnError(4002, "视频地址不能为空")
		v.ServeJSON()
		return
	}
	err := models.SaveVideo(title, subTitle, channelId, regionId, typeId, playUrl, uid, aliyunVideoId)
	if err == nil {
		v.Data["json"] = ReturnSuccess(0, "success", nil, 1)
		v.ServeJSON()
	} else {
		v.Data["json"] = ReturnError(5000, err)
		v.ServeJSON()
		return
	}
}
