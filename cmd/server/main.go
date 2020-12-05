package main

import (
	"github.com/SpaceXLaunchBot/site-backend/internal/api"
	"github.com/SpaceXLaunchBot/site-backend/internal/config"
	"github.com/SpaceXLaunchBot/site-backend/internal/database"
	"github.com/SpaceXLaunchBot/site-backend/internal/discord"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
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

	d := discord.NewClient(time.Second*10, time.Second*10)
	a := api.NewApi(db, d)
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/guildswithsubscribed", a.GuildsWithSubscribed).Methods("GET")
	r.HandleFunc("/api/updatesubscribedchannel", a.UpdateSubscribedChannel).Methods("POST")

	// Make sure the working directory has /static in it!
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Println("Serving http on all available interfaces @ port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
