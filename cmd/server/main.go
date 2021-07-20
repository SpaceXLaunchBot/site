package main

import (
	"github.com/SpaceXLaunchBot/site/internal/api"
	"github.com/SpaceXLaunchBot/site/internal/config"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

// TODO: Use Gin's logging instead of log.whatever?

func main() {
	// Assume that our dev environment in always Windows and production is always not.
	host := "spacexlaunchbot.dev"
	proto := "https:"
	port := ""
	if runtime.GOOS == "windows" {
		host = "localhost"
		proto = "http:"
		port = ":8080"
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	baseUrl := proto + "//" + host + port
	log.Printf("Base URL: %s", baseUrl)

	c, err := config.Get()
	if err != nil {
		log.Fatalf("config.Get error: %s", err)
	}

	db, err := database.NewDb(c)
	if err != nil {
		log.Fatalf("database.NewDb error: %s", err)
	}

	d := discord.NewClient(c.OAuthClientId, c.OAuthClientSecret, baseUrl+"/login")
	a := api.NewApi(db, d, host, proto)

	router := gin.Default()

	// https://fantashit.com/inability-to-use-for-static-files/
	router.Use(static.Serve("/", static.LocalFile("./frontend_build", false)))

	routerApi := router.Group("/api")
	routerApiAuthorized := routerApi.Group("/", a.SessionMiddleware())
	routerApiWithGuilds := routerApiAuthorized.Group("/", a.GuildListMiddleware())

	routerApiWithGuilds.GET("/subscribed", a.SubscribedChannels)
	routerApiWithGuilds.DELETE("/channel", a.DeleteChannel)
	routerApiWithGuilds.PUT("/channel", a.UpdateChannel)

	routerApi.GET("/stats", a.Stats)
	routerApiAuthorized.GET("/userinfo", a.UserInfo)

	routerApiAuthorized.GET("/auth/logout", a.HandleDiscordLogout)
	routerApiAuthorized.GET("/auth/verify", a.VerifySession)

	router.GET("/login", a.HandleDiscordLogin)

	// Due to React Router we have these routes that should all just server the index file.
	indexPath := "./frontend_build/index.html"
	router.StaticFile("/", indexPath)
	router.StaticFile("/commands", indexPath)
	router.StaticFile("/settings", indexPath)
	router.StaticFile("/stats", indexPath)

	log.Fatal(router.Run())
}
