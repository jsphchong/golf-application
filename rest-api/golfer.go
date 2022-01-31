package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateGolfer(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var g Golfer
	json.Unmarshal(body, &g)


}