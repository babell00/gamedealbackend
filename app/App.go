package app

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/babell00/gamedealbackend/game"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(){
	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/gamesearch", game.FindGames).Methods("GET")
}

func (a *App) Run(port string){
	log.Printf("Listening on port : %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), a.Router)

	if err != nil {
		log.Fatal(err)
	}
}
