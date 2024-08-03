package controllers

import (
	"encoding/json"
	"net/http"
	"server/structs"
)

func CreateQuestionOption(w http.ResponseWriter, r *http.Request) {
	var q structs.QuestionOption

	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
