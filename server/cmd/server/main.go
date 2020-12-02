package main

import (
	"github.com/gorilla/mux"
	"github.com/psidex/SpaceXLaunchBotSite/internal/api"
	"github.com/psidex/SpaceXLaunchBotSite/internal/config"
	"github.com/psidex/SpaceXLaunchBotSite/internal/database"
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

	a := api.NewApi(db)
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/guildswithsubscribed", a.GuildsWithSubscribed).Methods("GET")
	r.HandleFunc("/api/updatesubscribedchannel", a.UpdateSubscribedChannel).Methods("POST")

	// Make sure the working directory has /static in it!
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Println("Serving http on all available interfaces @ port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
