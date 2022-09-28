package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Ant0nYurchenk0/TutorialGoProject/pkg/config"
	"github.com/Ant0nYurchenk0/TutorialGoProject/pkg/handlers"
	"github.com/Ant0nYurchenk0/TutorialGoProject/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	render.NewTemplates(&app)

	app.TemplateCahce = tc
	app.UseCahce = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	log.Printf(fmt.Sprintf("starting at port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
