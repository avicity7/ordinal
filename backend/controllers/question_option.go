package controllers

import (
	"encoding/json"
	"net/http"
	"server/services"
	"server/structs"
)

func CreateQuestionOption(w http.ResponseWriter, r *http.Request) {
	var q structs.QuestionOption

	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.CreateQuestionOption(q.QuestionID, q.Body, q.Order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
