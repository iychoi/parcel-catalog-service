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
	"encoding/json"
	"net/http"

	"github.com/iychoi/parcel-catalog-service/pkg/database"
)

func (service *Service) listDatasetsHandler(writer http.ResponseWriter, request *http.Request) {
	database.CreateDB()

	datasets, err := database.GetAllDatasets()
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	json.NewEncoder(writer).Encode(datasets)
}
