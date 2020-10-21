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

package dataset

import (
	"encoding/json"
)

// Dataset contains information about a dataset
// Tags can have folowing additional data
// - citation
// - doi
// - publication year
type Dataset struct {
	ID          int64             `json:"id"`
	Name        string            `json:"name"` // contain space
	Creator     string            `json:"creator"`
	Description string            `json:"description"`
	URL         string            `json:"url"`  // e.g., irods://xxx or https://
	Host        string            `json:"host"` // e.g., CyVerse
	Rights      string            `json:"rights"`
	Tags        map[string]string `json:"tags"`
}

// Objectify returns Dataset from json byte array
func Objectify(jsonBytes []byte) *Dataset {
	var dataset Dataset
	json.Unmarshal(jsonBytes, &dataset)
	return &dataset
}

// Stringify returns string from Dataset
func Stringify(dataset *Dataset) string {
	jsonBytes, _ := json.Marshal(dataset)
	return string(jsonBytes)
}
