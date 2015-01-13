package server

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/inconshreveable/olive"
	"net/http"
	"os"

	log "gopkg.in/inconshreveable/log15.v2"
)

func RunTodayServer(bindAddr string) {
	fmt.Printf("Starting server on http://%s\n", bindAddr)
	exit(runAPI(bindAddr))
}

func exit(err error) {
	if err == nil {
		os.Exit(0)
	} else {
		log.Crit("command failed", "err", err)
		os.Exit(1)
	}
}

func runAPI(bindAddr string) error {
	handler := apiService()
	return http.ListenAndServe(bindAddr, handler)
}

func apiService() http.Handler {
	m := martini.NewRouter()

	return buildApiService(m)
}

func buildApiService(router martini.Router) http.Handler {
	o := olive.New(router)
	o.Debug = true

	//o.Get("/", o.Endpoint(getRootPage()))

	// Static file serving
	//o.Get("/", o.Endpoint(getRootPage()))
	router.Get("/bundle.css", func() string { return "body { }" })
	router.Get("/bundle.js", func() string { return "console.log('hello');" })

	m := martini.New()

	m.Use(martini.Static("public"))
	m.Action(router.Handle)

	return m
}

func getRootPage() martini.Handler {
	return func(r olive.Response, req *http.Request, params martini.Params) {
		r.Encode([]string{"root", "page"})
	}
}

func getBundledCSS() {
}

func getBundledJS() {
}
