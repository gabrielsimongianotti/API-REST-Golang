package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return session
}

var collection = getSession().DB("curso_go").C("movies")

var movies = Movies{
	Movie{"batmen", 2017, "super heroe"},
	Movie{"batmen1", 2017, "ola"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ola mundo ok.")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var result []Movie
	err := collection.Find(nil).Sort("-_id").All(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("resultado ", result)
	}
	json.NewEncoder(w).Encode(result)
	// fmt.Fprintf(w, "filme")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(movie_id)
	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	response(w, 200, results)
}

func response(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contato")
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err2 := collection.Insert(movie_data)

	if err2 != nil {
		w.WriteHeader(500)
		return
	}

	response(w, 200, movie_data)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": movie_data}
	err = collection.Update(document, change)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	response(w, 200, movie_data)
}

type Message struct {
	Status  string `json :"status"`
	Message string `json :"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

func MovieRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	message := new(Message)

	message.setStatus("seccess")
	message.setMessage("delete ID " + movie_id)

	results := message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
