# couchdb
Lightweight CouchDB interface for [go lang](https://golang.org/).

![alt text][couchdb logo] ![alt text][golang logo]

Go get the library `go get https://github.com/liamitus/couchdb` and you're good to go:
-------------------------------------------------------------------------------

```golang
import (
    "encoding/json"
    "fmt"
    "io/ioutil"

    "github.com/liamitus/couchdb"
)

var addr string = "http://127.0.0.1:5984"

// Check that CouchDB is running/reachable (will panic if not).
couchdb.Open(addr)

// Open a connection to the 'accounts' database.
// Will create the database if it doesn't exist.
var db couchdb.Database = couchdb.Open(fmt.Sprintf("%s/accounts", addr))

// Here's a sample struct for the purpose of this example.
type Account struct {
    Id  string `json:"_id,omitempty"`
    Rev string `json:"_rev,omitempty"`
    
    // Specifies the name of the account holder.
    Name string
}

// Create a new account in the 'accounts' database.
id := "1" // You'll probably want a UUID here.
account := Account{ Name: "Justice Beaver" }
b, _ := json.Marshal(account)
resp, err := db.Put(id, b)
if err != nil || resp.StatusCode != 201 {
    panic(fmt.Sprintf("Error creating %v\n%v", account, resp))
}

// Retrieve the account.
resp, err := db.Get(id, nil)
if err != nil {
    panic(err)
}
body, _ := ioutil.ReadAll(resp.Body)
var retrieved Account
json.Unmarshal(body, &retrieved)

// Delete the account.
url := fmt.Sprintf("%s?rev=%s", id, retrieved.Rev)
resp, err := db.Delete(url, nil)
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

[couchdb logo]: http://couchdb.apache.org/image/couch.png "Couch DB rocks!"
[golang logo]: https://golang.org/doc/gopher/frontpage.png "Go rocks too!"
