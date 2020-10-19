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
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Service object contains configuration parameters
type Service struct {
	config *Config
	
	server *Server
}

// NewService returns new service
func NewService(conf *Config) *Service {
	router := mux.NewRouter().StrictSlash(true),

	server := http.Server{
		Addr: MakeServiceAddress(conf),
		Handler: router,
	}

	svc := &Service{
		config: conf,
		server: server,
	}

	router.handleFunc(conf.RequestRoot, svc.handlerFunc)

	return svc
}

// Run runs the service
func (service *Service) Run() error {
	log.Printf("Listening for connections on address: %#v", addr)
	err := service.server.ListenAndServe()
	log.Printf("Stopped listening for connections on address: %#v", addr)
	
	return err
}
