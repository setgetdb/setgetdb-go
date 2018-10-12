package main

import (
	"encoding/json"
	. "github.com/setgetdb/setgetdb/setgetdb"
	"io/ioutil"
	"net/http"
	"os"
)

type ServerSetGetDb struct {
	db *Database
}

type keyValue struct {
	Key string `json:"key"`
	Value string `json:"value"`
}


func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	dbName := getenv("DB_NAME", "DATABASE")
	port := getenv("PORT", "10101")
	db := NewDatabase(dbName)
	server := ServerSetGetDb{db}
	http.HandleFunc("/set", server.Set)
	http.HandleFunc("/get", server.Get)
	http.HandleFunc("/delete", server.Delete)
	http.ListenAndServe(":" + port, nil)
}

func (server *ServerSetGetDb) Set(w http.ResponseWriter, r *http.Request) {
	buffer, _ := ioutil.ReadAll(r.Body)
	var request keyValue
	json.Unmarshal(buffer, &request)
	server.db.Set(request.Key, request.Value)
	w.Write(buffer)
}

func (server *ServerSetGetDb) Get(w http.ResponseWriter, r *http.Request) {
	buffer, _ := ioutil.ReadAll(r.Body)
	var request keyValue
	json.Unmarshal(buffer, &request)
	_, value := server.db.Get(request.Key)
	result := keyValue{ Key: request.Key, Value: value }
	resultJson, _ := json.Marshal(result)
	w.Write(resultJson)
}

func (server *ServerSetGetDb) Delete(w http.ResponseWriter, r *http.Request) {
	buffer, _ := ioutil.ReadAll(r.Body)
	var request keyValue
	json.Unmarshal(buffer, &request)
	server.db.Delete(request.Key)
	result := keyValue{ Key: request.Key }
	resultJson, _ := json.Marshal(result)
	w.Write(resultJson)
}
