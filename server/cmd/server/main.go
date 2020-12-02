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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/guildswithsubscribed", a.GuildsWithSubscribed)

	// Make sure the working directory has /static in it!
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	log.Println("Serving http on all available interfaces @ port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
