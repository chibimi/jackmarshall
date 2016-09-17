package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewUpdateScenarioHandler(db *mgo.Session) httprouter.Handle {
	collection := db.DB("jackmarshall").C("scenario")

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		id := p.ByName("id")

		var scenario Scenario
		err := json.NewDecoder(r.Body).Decode(&scenario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = collection.UpdateId(bson.ObjectIdHex(id), &scenario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}