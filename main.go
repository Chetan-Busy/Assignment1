package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type company struct {
	Name         string    `json:"name"`
	AgeInYears   float64   `json:"age_in_years"`
	Origin       string    `json:"origin"`
	HeadOffice   string    `json:"head_office"`
	Address      []address `json:"address"`
	Sponsers     sponser   `json:"sponsers"`
	Revenue      string    `json:"revenue"`
	NoOfEmployee float64   `json:"no_of_employee"`
	StrText      []string  `json:"str_text"`
	IntText      []float64 `json:"int_text"`
}

type address struct {
	Street   string  `json:"street"`
	Landmark string  `json:"landmark"`
	City     string  `json:"city"`
	Pincode  float64 `json:"pincode"`
	State    string  `json:"state"`
}

type sponser struct {
	Name string `json:"name"`
}

func setKeyValue(key string, val interface{}, src map[string]interface{}) {
	if _, ok := src[key]; ok {
		src[key] = val
		return
	}
	for _, v := range src {
		rVal := reflect.ValueOf(v)
		switch rVal.Kind() {
		case reflect.Map:
			setKeyValue(key, val, v.(map[string]interface{}))
		case reflect.Slice:
			for _, m := range v.([]interface{}) {
				if reflect.ValueOf(m).Kind() == reflect.Map {
					setKeyValue(key, val, m.(map[string]interface{}))
				}
			}
		}
	}
}

func removeKey(key string, src map[string]interface{}) {
	if _, ok := src[key]; ok {
		delete(src, key)
		return
	}
	for _, v := range src {
		rVal := reflect.ValueOf(v)
		switch rVal.Kind() {
		case reflect.Map:
			removeKey(key, v.(map[string]interface{}))
		case reflect.Slice:
			for _, m := range v.([]interface{}) {
				if reflect.ValueOf(m).Kind() == reflect.Map {
					removeKey(key, m.(map[string]interface{}))
				}
			}
		}
	}
}

func main() {
	var inp string = `{
		"name" : "Tolexo Online Pvt. Ltd",
		"age_in_years" : 8.5,
		"origin" : "Noida",
		"head_office" : "Noida, Uttar Pradesh",
		"address" : [
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			},
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				
				"state" : "Uttar Pradesh"
			}
		],
		"sponsers" : {
			"name" : "One"
		},
		"revenue" : "19.8 million$",
		"no_of_employee" : 630,
		"str_text" : ["one","two"],
		"int_text" : [1,3,4],
		"city": "abc"
	}`

	var mp map[string]interface{}

	err := json.Unmarshal([]byte(inp), &mp)

	if err != nil {
		panic(err)
	}

	fmt.Println(mp)

	city := "New Delhi"
	key := "city"

	setKeyValue(key, city, mp)

	fmt.Println(mp)

	removeKey(key, mp)

	fmt.Println(mp)
}
