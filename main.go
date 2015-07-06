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
func Open(url string) Database {
	return Database{url}
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
