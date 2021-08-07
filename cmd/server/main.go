package main

import (
	"log"
	"runtime"

	"github.com/SpaceXLaunchBot/site/internal/api"
	"github.com/SpaceXLaunchBot/site/internal/config"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Assume that our dev environment in always Windows and production is always not.
const inDev bool = runtime.GOOS == "windows"

// devString returns one of the given strings depending on if we are in development mode or not.
func devString(devVar string, notDevVar string) string {
	if inDev {
		return devVar
	}
	return notDevVar
}

func main() {
	if inDev {
		log.Println("Running in development mode")
	}

	gin.SetMode(devString(gin.DebugMode, gin.ReleaseMode))

	staticDir := devString("./frontend/build", "./frontend_build")

	host := devString("localhost", "spacexlaunchbot.dev")
	proto := devString("http:", "https:")
	port := devString(":8080", "")
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
	router.Use(static.Serve("/", static.LocalFile(staticDir, false)))

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
	indexPath := staticDir + "/index.html"
	router.StaticFile("/", indexPath)
	router.StaticFile("/commands", indexPath)
	router.StaticFile("/settings", indexPath)
	router.StaticFile("/stats", indexPath)

	log.Println("Starting Gin")
	log.Fatal(router.Run())
}
