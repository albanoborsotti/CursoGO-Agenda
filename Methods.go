package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola esto es una agenda")
}

//Valida por medio de expresiones regulares el email
func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func responseContact(w http.ResponseWriter, status int, results Contact) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)

}
func responseContacts(w http.ResponseWriter, status int, results []Contact) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)

}

var collection = getSession().DB("CursoGo").C("contact")

func ContactAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var contactData Contact
	err := decoder.Decode(&contactData)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()
	err = collection.Insert(contactData)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println(contactData.Email)

	if validateEmail(contactData.Email) {
		responseContact(w, 200, contactData)
	} else {
		//fmt.Fprinl(w, "El Email ingresado no es valido.")
		fmt.Println("El Email ingresado no es valido.")
		w.WriteHeader(404)
	}

}
func ContactList(w http.ResponseWriter, r *http.Request) {

	var results []Contact
	err := collection.Find(nil).Sort("lastname").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}
	responseContacts(w, 200, results)
}

func ContactUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contactID := params["id"]

	if !bson.IsObjectIdHex(contactID) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(contactID)

	decoder := json.NewDecoder(r.Body)

	var contactData Contact
	err := decoder.Decode(&contactData)

	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}
	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": contactData}

	err = collection.Update(document, change)

	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}

	responseContact(w, 200, contactData)

}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ContactDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contactID := params["id"]

	if !bson.IsObjectIdHex(contactID) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(contactID)

	err := collection.RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	results := Message{"Success", "El contacto con ID " + contactID + " ha sido borrado con exito"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)

}

func ContactSearch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contact := params["txt"]
	fmt.Println(contact)
	var results []Contact
	err := collection.Find(bson.M{"name": contact}).All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}
	responseContacts(w, 200, results)
}
