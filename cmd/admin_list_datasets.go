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

	"github.com/iychoi/parcel-catalog-service/pkg/database"
	"github.com/iychoi/parcel-catalog-service/pkg/dataset"
	"github.com/iychoi/parcel-catalog-service/pkg/service"
)

func main() {
	var version bool

	// Parse parameters
	flag.BoolVar(&version, "version", false, "Print service version information")

	flag.Parse()

	// Handle Version
	if version {
		info, err := service.GetVersionJSON()

		if err != nil {
			log.Fatal(err)
		}

		log.Println(info)
		os.Exit(0)
	}

	datasets, err := database.GetAllDatasets()

	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, ds := range datasets {
		log.Printf(dataset.Stringify(ds))
		count++
	}

	log.Printf("Displayed %d datasets\n", count)

	os.Exit(0)
}
