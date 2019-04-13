package main

import (
	"encoding/json"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func jsonify(d interface{}) string {
	s, err := json.Marshal(d)
	check(err)
	return string(s)
}
