package couchdb

import (
	"bytes"
	"fmt"
	"net/http"
)

// The database itself.
type Database struct {
	Url string
}

// Constructor for a new Database.
// Accepts optional failure and success messages (in that order).
// Will panic if the database is not running or not reachable.
func Open(url string, msgs ...string) Database {
	db := Database{url}
	// Panic if the database is not running.
	if !db.exists("") {
		panic(fmt.Sprintf("DB is not running at %q", db.Url))
	}
	// Otherwise create the table does not exist.
	if !db.exists(url, msgs...) {
		db.Put("", nil)
	}
	return db
}

// Perform a GET request to the database.
func (d *Database) Get(path string, data []byte) (*http.Response, error) {
	return d.query("GET", path, data)
}

// Perform a PUT request to the database.
func (d *Database) Put(path string, data []byte) (*http.Response, error) {
	return d.query("PUT", path, data)
}

// Perform a DELETE request to the database.
func (d *Database) Delete(path string, data []byte) (*http.Response, error) {
	return d.query("DELETE", path, data)
}

// Simplifies making a request to the database into a single function call.
func (d *Database) query(requestType string, path string, data []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", d.Url, path)
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}

// Checks a given URL to see if the response returns without error.
// Accepts optional failure and success messages (in that order).
func (d *Database) exists(url string, msgs ...string) bool {
	if resp, err := d.Get(url, nil); err != nil || resp.StatusCode != 200 {
		if len(msgs) > 0 {
			fmt.Println(msgs[0])
		}
		if err != nil {
			fmt.Println(err)
		}
		return false
	}

	if len(msgs) > 1 {
		fmt.Println(msgs[1])
	}

	return true
}
