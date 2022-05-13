package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["movie/backend/controllers:BarrageController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/barrage/save`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:BarrageController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:BarrageController"],
        beego.ControllerComments{
            Method: "BarrageWs",
            Router: `/barrage/ws`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:BaseController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelRegion",
            Router: `/channel/region`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:BaseController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: `/channel/type`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:CommentController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:CommentController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/comment/list`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:CommentController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Save",
            Router: `/comment/save`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:TopController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:TopController"],
        beego.ControllerComments{
            Method: "ChannelTop",
            Router: `/channel/top`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:TopController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:TopController"],
        beego.ControllerComments{
            Method: "TypeTop",
            Router: `/type/top`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginDo",
            Router: `/login/do`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "SaveRegister",
            Router: `/register/save`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:UserController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "SendMessageDo",
            Router: `/send/message`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelAdvert",
            Router: `/channel/advert`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelHotList",
            Router: `/channel/hot`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelRecommendRegionlist",
            Router: `/channel/recommend/region`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelRecommendTypeList",
            Router: `/channel/recommend/type`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelVideo",
            Router: `/channel/video`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "UserVideo",
            Router: `/user/video`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoEpisodesList",
            Router: `/video/episodes/list`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoInfo",
            Router: `/video/info`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["movie/backend/controllers:VideoController"] = append(beego.GlobalControllerRouter["movie/backend/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoSave",
            Router: `/video/save`,
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
