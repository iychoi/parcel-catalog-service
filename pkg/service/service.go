/*
Copyright 2020 CyVerse
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Service object contains configuration parameters
type Service struct {
	config *Config

	server *http.Server
}

// NewService returns new service
func NewService(conf *Config) *Service {
	router := mux.NewRouter().StrictSlash(true)

	server := &http.Server{
		Addr:    MakeServiceAddress(conf),
		Handler: router,
	}

	svc := &Service{
		config: conf,
		server: server,
	}

	// add handlers
	svc.addHandlers(router)

	return svc
}

// Run runs the service
func (service *Service) Run() error {
	log.Printf("Listening for connections on address: %#v", service.server.Addr)
	err := service.server.ListenAndServe()
	log.Printf("Stopped listening for connections on address: %#v", service.server.Addr)

	return err
}

// Handler handles requests
func (service *Service) rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Parcel-Catalog-Service Root Page\n")
}

// addHandlers adds all request Handlers
func (service *Service) addHandlers(router *mux.Router) {
	router.HandleFunc(service.config.RequestRoot, service.rootHandler)

	router.HandleFunc(makeRequestPath(service.config.RequestRoot, "/datasets"), service.listDatasetsHandler).Methods("GET")
}

func makeRequestPath(requestRoot string, path string) string {
	if strings.HasSuffix(requestRoot, "/") && strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s%s", strings.TrimRight(requestRoot, "/"), path)
	} else if !strings.HasSuffix(requestRoot, "/") && !strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s/%s", requestRoot, path)
	} else {
		return fmt.Sprintf("%s%s", requestRoot, path)
	}
}
