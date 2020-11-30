package main

import (
    "github.com/gorilla/mux"
    "github.com/psidex/SpaceXLaunchBotSite/internal/api"
    "log"
    "net/http"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/login", api.Login)
    // Make sure the working directory has /static in it!
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

    log.Println("Serving http on all available interfaces @ port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
