package main

import (
	"github.com/SpaceXLaunchBot/site/internal/api"
	"github.com/SpaceXLaunchBot/site/internal/config"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend_build/index.html")
}

func main() {
	host := "spacexlaunchbot.dev"
	proto := "https:"
	port := ""
	if runtime.GOOS == "windows" {
		host = "localhost"
		proto = "http:"
		port = ":8080"
	}

	c, err := config.Get()
	if err != nil {
		log.Fatalf("config.Get error: %s", err)
	}

	db, err := database.NewDb(c)
	if err != nil {
		log.Fatalf("database.NewDb error: %s", err)
	}

	d := discord.NewClient(c.ClientId, c.ClientSecret, proto+"//"+host+port+"/login")
	a := api.NewApi(db, d, host, proto)
	r := mux.NewRouter().StrictSlash(true)

	rApi := r.PathPrefix("/api").Subrouter()
	// Routes under rApiSession get passed context about the users current session.
	rApiSession := r.PathPrefix("/api").Subrouter()
	rApiSession.Use(a.SessionMiddleware)

	rApiSession.HandleFunc("/subscribed", a.SubscribedChannels).Methods("GET")
	rApiSession.HandleFunc("/channel", a.DeleteChannel).Methods("DELETE")
	rApiSession.HandleFunc("/channel", a.UpdateChannel).Methods("PUT")

	rApi.HandleFunc("/stats", a.Stats).Methods("GET")
	rApiSession.HandleFunc("/userinfo", a.UserInfo).Methods("GET")

	rApi.HandleFunc("/auth/login", a.HandleDiscordLogin).Methods("GET")
	rApiSession.HandleFunc("/auth/logout", a.HandleDiscordLogout).Methods("GET")
	rApiSession.HandleFunc("/auth/verify", a.VerifySession).Methods("GET")

	// Due to React Router we have these routes that should all just server the index file.
	r.HandleFunc("/", serveIndex)
	r.HandleFunc("/commands", serveIndex)
	r.HandleFunc("/settings", serveIndex)
	r.HandleFunc("/stats", serveIndex)
	r.HandleFunc("/login", serveIndex)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend_build")))

	log.Println("Serving http on all available interfaces @ port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
