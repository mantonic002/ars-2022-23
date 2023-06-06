// Config API
//
//	Title: Config API
//
//	Schemes: http
//	Version: 0.0.1
//	BasePath: /
//
//	Produces:
//	  - application/json
//
// swagger:meta
package main

import (
	"ars-2022-23/ConfigStore"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	store, err := ConfigStore.New()
	if err != nil {
		log.Fatal(err)
	}

	groupStore, err := ConfigStore.NewGroup()
	if err != nil {
		log.Fatal(err)
	}

	server := configServer{
		store:      store,
		groupStore: groupStore,
	}
	router.HandleFunc("/config/", count(server.createConfigHandler)).Methods("POST")
	router.HandleFunc("/configs/", count(server.getAllHandler)).Methods("GET")
	router.HandleFunc("/config/{id}/", count(server.getConfigHandler)).Methods("GET")
	router.HandleFunc("/configs/{labels}", count(server.getConfigByLabelHandler)).Methods("GET")

	router.HandleFunc("/config/{id}/", count(server.delConfigHandler)).Methods("DELETE")

	router.HandleFunc("/group/", count(server.createGroupHandler)).Methods("POST")
	router.HandleFunc("/group/{groupId}/config/{id}/", count(server.addConfigToGroup)).Methods("PUT")

	router.HandleFunc("/groups/", count(server.getAllGroupsHandler)).Methods("GET")
	router.HandleFunc("/group/{id}/", count(server.getGroupHandler)).Methods("GET")

	router.HandleFunc("/group/{id}/", count(server.delGroupHandler)).Methods("DELETE")
	router.HandleFunc("/group/{groupId}/config/{id}/", count(server.delConfigFromGroupHandler)).Methods("DELETE")

	router.HandleFunc("/swagger.yaml", server.swaggerHandler).Methods("GET")
	router.Path("/metrics").Handler(metricsHandler())

	// SwaggerUI
	optionsDevelopers := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	developerDocumentationHandler := middleware.SwaggerUI(optionsDevelopers, nil)
	router.Handle("/docs", developerDocumentationHandler)

	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
