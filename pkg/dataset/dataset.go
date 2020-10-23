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
	"fmt"
	"strconv"
	"strings"
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

// PrintDataset prints dataset
func (ds *Dataset) PrintDataset(short bool, maxDescriptionLen int) {
	if short {
		fmt.Printf("[%d] %s\n", ds.ID, ds.Name)
		fmt.Printf("  Creator     : %s\n", ds.Creator)
		fmt.Printf("  Rights      : %s\n", ds.Rights)
		if len(ds.Description) > maxDescriptionLen {
			fmt.Printf("  Description : %s...[more]\n", ds.Description[:maxDescriptionLen])
		} else {
			fmt.Printf("  Description : %s\n", ds.Description)
		}
	} else {
		fmt.Printf("[%d] %s\n", ds.ID, ds.Name)
		fmt.Printf("  Name        : %s\n", ds.Name)
		fmt.Printf("  Creator     : %s\n", ds.Creator)
		fmt.Printf("  Host        : %s\n", ds.Host)
		fmt.Printf("  Description : %s\n", ds.Description)
		fmt.Printf("  Rights      : %s\n", ds.Rights)
		fmt.Printf("  URL         : %s\n", ds.URL)
		for k, v := range ds.Tags {
			fmt.Printf("  %-12s: %s\n", k, v)
		}
	}
}

// ContainsKeywords checks if the dataset has the given keywords
func (ds *Dataset) ContainsKeywords(keywords []string) bool {
	for _, keyword := range keywords {
		if keyword == strconv.FormatInt(ds.ID, 10) {
			return true
		}

		if strings.Contains(strings.ToLower(ds.Name), strings.ToLower(keyword)) {
			return true
		}

		if strings.Contains(strings.ToLower(ds.Creator), strings.ToLower(keyword)) {
			return true
		}

		if strings.Contains(strings.ToLower(ds.Host), strings.ToLower(keyword)) {
			return true
		}

		if strings.Contains(strings.ToLower(ds.Description), strings.ToLower(keyword)) {
			return true
		}

		if strings.Contains(strings.ToLower(ds.Rights), strings.ToLower(keyword)) {
			return true
		}

		if strings.Contains(strings.ToLower(ds.URL), strings.ToLower(keyword)) {
			return true
		}

		for _, v := range ds.Tags {
			if strings.Contains(strings.ToLower(v), strings.ToLower(keyword)) {
				return true
			}
		}
	}
	return false
}

// Objectify returns Dataset from json byte array
func Objectify(jsonBytes []byte) *Dataset {
	var ds Dataset
	json.Unmarshal(jsonBytes, &ds)
	return &ds
}

// Listify returns Dataset array from json byte array
func Listify(jsonBytes []byte) []*Dataset {
	var ds []Dataset
	json.Unmarshal(jsonBytes, &ds)

	dsp := []*Dataset{}
	for _, d := range ds {
		dsp = append(dsp, &d)
	}

	return dsp
}

// Stringify returns string from Dataset
func Stringify(dataset *Dataset) string {
	jsonBytes, _ := json.Marshal(dataset)
	return string(jsonBytes)
}
