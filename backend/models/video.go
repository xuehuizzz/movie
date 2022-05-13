package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            time.Time
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime time.Time
	Comment            int
	UserId             int
}
type VideoData struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       time.Time
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
}

type Episodes struct {
	Id      int
	Title   string
	AddTime time.Time
	Num     int
	PlayUrl string
	Comment int
}

func init() {
	orm.RegisterModel(new(Video))
}
func GetChannelHotList(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id, title, sub_title, add_time, img, img1,episodes_count,is_end from video where status=1 "+
		"and is_hot=1 and channel_id=? order by episodes_count desc limit 9", channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelRecommendRegionList(channelId int, regionId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id,title,sub_title,add_time,img,img1,episodes_count,is_end from video where status=1 "+
		"and is_recommend=1 and region_id=? and channel_id=? order by episodes_update_time desc limit 9", regionId, channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelRecommendTypeList(channelId int, typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id,title,sub_title,add_time,img,img1,episodes_count,is_end from video where status=1 "+
		"and is_recommend=1 and type_id=? and channel_id=? order by episodes_update_time desc limit 9", typeId, channelId).QueryRows(&videos)
	return num, videos, err
}

func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params
	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" {
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		qs = qs.Filter("is_end", 1)
	}
	if sort == "episodesUpdateTime" {
		qs = qs.OrderBy("-episodes_update_time")
	} else if sort == "comment" {
		qs = qs.OrderBy("-comment")
	} else if sort == "addTime" {
		qs = qs.OrderBy("-add_time")
	} else {
		qs = qs.OrderBy("-add_time")
	}
	nums, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit, offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	return nums, videos, err
}

func GetUserVideo(uid int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id,title,sub_title,img,img1,add_time,episodes_count, is_end from video where user_id=? order by add_time desc", uid).QueryRows(&videos)
	return num, videos, err
}

func GetVideoInfo(videoId int) (Video, error) {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("select * from video where id=? limit 1", videoId).QueryRow(&video)
	return video, err
}

func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.Raw("select id,title,add_time,num,play_url,comment from video_episodes where video_id=? and status=1 order by num asc", videoId).QueryRows(&episodes)
	return num, episodes, err
}

func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id,title,sub_title,img,img1,add_time,episodes_count,is_end from video where status=1 and channel_id=? order by comment limit 10", channelId).QueryRows(&videos)
	return num, videos, err
}

func GetTypeTop(typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("select id,title,sub_title,img,img1,add_time,episodes_count,is_end from video where status=1 and type_id=? order by comment limit 10", typeId).QueryRows(&videos)
	return num, videos, err
}

func SaveVideo(title string, subTitle string, channelId int, regionId int, typeId int, playUrl string, userId int, aliyunVideoId string) error {
	o := orm.NewOrm()
	var video Video
	video.Title = title
	video.SubTitle = subTitle
	video.AddTime = time.Now()
	video.Img = ""
	video.Img1 = ""
	video.EpisodesCount = 1
	video.IsEnd = 1
	video.ChannelId = channelId
	video.RegionId = regionId
	video.Status = 1
	video.TypeId = typeId
	video.EpisodesUpdateTime = time.Now()
	video.Comment = 0
	video.UserId = userId
	videoId, err := o.Insert(&video)
	if err == nil {
		if aliyunVideoId != "" {
			playUrl = ""
		}
		o.Raw("insert into video_episodes (title,add_time,num,video_id,play_url,status,comment,aliyun_video_id) values (?,?,?,?,?,?,?,?)", subTitle, time.Now(), 1, videoId, playUrl, 1, 0, aliyunVideoId).Exec()
	}
	return err
}


func SaveAliyunVideo(videoId string, log string) error {
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO aliyun_video (video_id, log, add_time) VALUES (?,?,?)", videoId, log, time.Now().Unix()).Exec()
	fmt.Println(err)
	return err
}

//获取所有视频数据
func GetAllList() (int64, []Video, error) {
	o := orm.NewOrm()
	var videos []Video
	num, err := o.Raw("SELECT id,title,sub_title,status,add_time, img,img1,channel_id,type_id,region_id,user_id,episodes_count,episodes_update_time,is_end,is_hot,is_recommend,comment FROM video").QueryRows(&videos)
	return num, videos, err
}
