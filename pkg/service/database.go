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
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Dataset struct {
	ID          int    `json:"id"`
	Name        string `json:"name"` // contain space
	Creator     string `json:"creator"`
	Description string `json:"description"`
	URL         string `json:"url"`  // e.g., irods://xxx or https://
	Host        string `json:"host"` // e.g., CyVerse
	Rights      string `json:"rights"`
	Tags        string `json:"tags"` // has json string in it
}

const (
	DB_FILE_NAME = "parcel.db"
)

// Tags can have folowing additional data
// - citation
// - doi
// - publication year

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// CreateDB creates a new database
func (service *Service) CreateDB() {
	if !fileExists(DB_FILE_NAME) {
		log.Println("file not exists")

		file, err := os.Create(DB_FILE_NAME)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		sqlite, _ := sql.Open("sqlite3", DB_FILE_NAME)
		defer sqlite.Close()

		// create a table
		sql := `CREATE TABLE dataset (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"name" varchar(50),
			"creator" varchar(50),
			"description" TEXT,
			"url" varchar(500),
			"host" varchar(50),
			"rights" varchar(500),
			"tags" TEXT
		);`

		statement, err := sqlite.Prepare(sql)
		if err != nil {
			log.Fatal(err)
		}

		_, err = statement.Exec()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Created 'dataset' table")
	}
}

// openDatabase opens database safely
func openDatabase() (*sql.DB, error) {
	sqlite, _ := sql.Open("sqlite3", DB_FILE_NAME)
	return sqlite, nil
}

// openDatabase closes database safely
func closeDatabase(sqlite *sql.DB) error {
	sqlite.Close()
	return nil
}

// GetAllDatasets returns all datasets
func (service *Service) GetAllDatasets() ([]*Dataset, error) {
	sqlite, err := openDatabase()
	if err != nil {
		return nil, err
	}

	defer closeDatabase(sqlite)

	sql := "SELECT id, name, creator, description, url, host, rights, tags FROM dataset;"
	statement, err := sqlite.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}

	datasets := []*Dataset{}
	for rows.Next() {
		var dataset Dataset

		err := rows.Scan(&dataset.ID, &dataset.Name, &dataset.Creator, &dataset.Description, &dataset.URL, &dataset.Host, &dataset.Rights, &dataset.Tags)
		if err != nil {
			return nil, err
		}

		datasets = append(datasets, &dataset)
	}

	return datasets, nil
}
