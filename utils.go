package main

import "encoding/json"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func jsonify(d interface{}) string {
	s, err := json.Marshal(d)
	check(err)
	return string(s)
}
