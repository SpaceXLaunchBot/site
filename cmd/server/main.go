package main

import (
	"github.com/SpaceXLaunchBot/site/internal/api"
	"github.com/SpaceXLaunchBot/site/internal/config"
	"github.com/SpaceXLaunchBot/site/internal/database"
	"github.com/SpaceXLaunchBot/site/internal/discord"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	c, err := config.Get()
	if err != nil {
		log.Fatalf("config.Get error: %s", err)
	}

	db, err := database.NewDb(c)
	if err != nil {
		log.Fatalf("database.NewDb error: %s", err)
	}

	d := discord.NewClient()
	a := api.NewApi(db, d)
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/subscribed", a.SubscribedChannels).Methods("GET")
	r.HandleFunc("/api/channel", a.DeleteChannel).Methods("DELETE")
	r.HandleFunc("/api/channel", a.UpdateChannel).Methods("PUT")
	r.HandleFunc("/api/metrics", a.Metrics).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend_build")))

	log.Println("Serving http on all available interfaces @ port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
