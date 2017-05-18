package core

import (
	"encoding/json"
	"reflect"
)

var base = reflect.TypeOf(simpleBasicStruct{})
var basePage = reflect.TypeOf(basePageStruct{})

// UserTypesMap handles object types for creating
// objects from incoming webhook calls from facebook
var UserTypesMap = map[string]reflect.Type{
	"about":       base,
	"activities":  basePage,
	"birthday":    base,
	"books":       basePage,
	"education":   reflect.TypeOf(education{}),
	"email":       base,
	"events":      reflect.TypeOf(events{}),
	"feed":        reflect.TypeOf(feed{}),
	"first_name":  base,
	"friends":     reflect.TypeOf(friends{}),
	"gender":      base,
	"hometown":    basePage,
	"last_name":   base,
	"likes":       basePage,
	"live_videos": reflect.TypeOf(live_videos{}),
	"locale":      base,
	"location":    basePage,
	"movies":      basePage,
	"music":       basePage,
	"name":        reflect.TypeOf(name{}),
	"photos":      reflect.TypeOf(photos{}),
	"pic":         base,
	"platform":    base,
	"quotes":      base,
	"religion":    base,
	"status":      reflect.TypeOf(status{}),
	"television":  basePage,
	"videos":      reflect.TypeOf(live_videos{}),
	"website":     base,
	"work":        reflect.TypeOf(work{}),
}

type simpleBasicStruct struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type basePageStruct struct {
	Field string `json:"field"`
	Verb  string `json:"verb"`
	Value struct {
		Page string `json:"page"`
	} `json:"value"`
}

type education struct {
	Type      string                `json:"type"`
	Education []educationExperience `json:"value"`
}

type events struct {
	Field string `json:"field"`
	Value struct {
		EventID string `json:"event_id"`
		Verb    string `json:"verb"`
	} `json:"value"`
}

type feed struct {
	Field string `json:"field"`
}

type friends struct {
	Field string `json:"field"`
	Value struct {
		UUID string `json:"uuid"`
	} `json:"value"`
	Verb string `json:"verb"`
}

type live_videos struct {
	Field string `json:"field"`
	Value struct {
		ID     int64  `json:"id"`
		Status string `json:"status"`
	} `json:"value"`
}

type name struct {
	Field          string          `json:"field"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	MiddleName     string          `json:"middle_name"`
	Value          string          `json:"value"`
	LocalizedNames json.RawMessage `json:"localized_names"`
}

type photos struct {
	Field string `json:"field"`
	Value struct {
		ObjectID string `json:"object_id"`
		Verb     string `json:"verb"`
	} `json:"value"`
}

type status struct {
	ID    int64  `json:"id"`
	Field string `json:"field"`
	Value string `json:"value"`
}

type work struct {
	Field string `json:"field"`
	Value []WorkExperience
}


