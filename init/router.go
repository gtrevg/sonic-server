package init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/cloudsonic/sonic-server/api"
	"github.com/cloudsonic/sonic-server/controllers"
	"github.com/twinj/uuid"
)

const requestidHeader = "X-Request-Id"

func init() {
	mapEndpoints()
	mapControllers()
	initFilters()
}

func mapEndpoints() {
	ns := beego.NewNamespace("/rest",
		beego.NSRouter("/ping.view", &api.SystemController{}, "*:Ping"),
		beego.NSRouter("/getLicense.view", &api.SystemController{}, "*:GetLicense"),

		beego.NSRouter("/getMusicFolders.view", &api.BrowsingController{}, "*:GetMusicFolders"),
		beego.NSRouter("/getIndexes.view", &api.BrowsingController{}, "*:GetIndexes"),
		beego.NSRouter("/getMusicDirectory.view", &api.BrowsingController{}, "*:GetMusicDirectory"),
		beego.NSRouter("/getSong.view", &api.BrowsingController{}, "*:GetSong"),
		beego.NSRouter("/getArtists.view", &api.BrowsingController{}, "*:GetArtists"),
		beego.NSRouter("/getArtist.view", &api.BrowsingController{}, "*:GetArtist"),
		beego.NSRouter("/getAlbum.view", &api.BrowsingController{}, "*:GetAlbum"),

		beego.NSRouter("/search2.view", &api.SearchingController{}, "*:Search2"),
		beego.NSRouter("/search3.view", &api.SearchingController{}, "*:Search3"),

		beego.NSRouter("/getCoverArt.view", &api.MediaRetrievalController{}, "*:GetCoverArt"),
		beego.NSRouter("/getAvatar.view", &api.MediaRetrievalController{}, "*:GetAvatar"),
		beego.NSRouter("/stream.view", &api.StreamController{}, "*:Stream"),
		beego.NSRouter("/download.view", &api.StreamController{}, "*:Download"),

		beego.NSRouter("/scrobble.view", &api.MediaAnnotationController{}, "*:Scrobble"),
		beego.NSRouter("/star.view", &api.MediaAnnotationController{}, "*:Star"),
		beego.NSRouter("/unstar.view", &api.MediaAnnotationController{}, "*:Unstar"),
		beego.NSRouter("/setRating.view", &api.MediaAnnotationController{}, "*:SetRating"),

		beego.NSRouter("/getAlbumList.view", &api.AlbumListController{}, "*:GetAlbumList"),
		beego.NSRouter("/getAlbumList2.view", &api.AlbumListController{}, "*:GetAlbumList2"),
		beego.NSRouter("/getStarred.view", &api.AlbumListController{}, "*:GetStarred"),
		beego.NSRouter("/getStarred2.view", &api.AlbumListController{}, "*:GetStarred2"),
		beego.NSRouter("/getNowPlaying.view", &api.AlbumListController{}, "*:GetNowPlaying"),
		beego.NSRouter("/getRandomSongs.view", &api.AlbumListController{}, "*:GetRandomSongs"),

		beego.NSRouter("/getPlaylists.view", &api.PlaylistsController{}, "*:GetPlaylists"),
		beego.NSRouter("/getPlaylist.view", &api.PlaylistsController{}, "*:GetPlaylist"),
		beego.NSRouter("/createPlaylist.view", &api.PlaylistsController{}, "*:CreatePlaylist"),
		beego.NSRouter("/updatePlaylist.view", &api.PlaylistsController{}, "*:UpdatePlaylist"),
		beego.NSRouter("/deletePlaylist.view", &api.PlaylistsController{}, "*:DeletePlaylist"),

		beego.NSRouter("/getUser.view", &api.UsersController{}, "*:GetUser"),
	)
	beego.AddNamespace(ns)

}

func mapControllers() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sync", &controllers.SyncController{})

	beego.ErrorController(&controllers.MainController{})
}

func initFilters() {
	var requestIdFilter = func(ctx *context.Context) {
		id := ctx.Input.Header(requestidHeader)
		if id == "" {
			id = uuid.NewV4().String()
		}
		ctx.Input.SetData("requestId", id)
	}

	var validateRequest = func(ctx *context.Context) {
		c := api.BaseAPIController{}
		// TODO Find a way to not depend on a controller being passed
		c.Ctx = ctx
		c.Data = make(map[interface{}]interface{})
		api.Validate(c)
	}

	beego.InsertFilter("/rest/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("/rest/*", beego.BeforeRouter, requestIdFilter)
	beego.InsertFilter("/rest/*", beego.BeforeRouter, validateRequest)
}
