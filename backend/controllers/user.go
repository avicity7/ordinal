package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/services"

	"github.com/patrickmn/go-cache"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var output interface{}
	value, found := config.Cache.Get("Users")
	if found {
		output = value
	} else {
		arr, err := services.GetUsers()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
		output = arr
		config.Cache.Set("Users", arr, cache.DefaultExpiration)
	}
	w.WriteHeader(200)
	formatted, _ := json.Marshal(output)
	w.Write(formatted)
}
