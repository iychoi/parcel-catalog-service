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

package main

import (
	"flag"
	"log"
	"os"

	"github.com/iychoi/parcel-catalog-service/pkg/service"
)

var (
	conf service.Config
)

func main() {
	var version bool

	// Parse parameters
	flag.StringVar(&conf.Address, "address", "", "service address")
	flag.IntVar(&conf.Port, "port", 80, "service port number")
	flag.StringVar(&conf.RequestRoot, "requestRoot", "/", "service request root dir")

	flag.BoolVar(&version, "version", false, "Print service version information")

	flag.Parse()

	// Handle Version
	if version {
		info, err := service.GetVersionJSON()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		log.Println(info)
		os.Exit(0)
	}

	log.Printf("Driver version: %s", service.GetServiceVersion())

	svc := service.NewService(&conf)
	if err := svc.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}
