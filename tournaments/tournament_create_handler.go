package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/julienschmidt/httprouter"
)

func NewCreateTournamentHandler(database *db.DB) httprouter.Handle {
	collection := database.Use("Tournaments")

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var tournament Tournament
		err := json.NewDecoder(r.Body).Decode(&tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(tournament)
		id, err := collection.Insert(map[string]interface{}{
			"name":       tournament.Name,
			"format":     tournament.Format,
			"slots":      tournament.Slots,
			"fee_amount": tournament.FeeAmount,
			"date":       tournament.Date,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(id)
	}
}