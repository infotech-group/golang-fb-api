package core

import "reflect"

// UserTypesMap handles object types for creating
// objects from incoming webhook calls from facebook
var UserTypesMap = map[string]reflect.Type{
	"about":      reflect.TypeOf(about{}),
	"activities": reflect.TypeOf(activities{}),
	"birthday":   reflect.TypeOf(birthday{}),
	"books":      reflect.TypeOf(books{}),
	"education":  reflect.TypeOf(education{}),
	"email":      reflect.TypeOf(email{}),
	"events":     reflect.TypeOf(events{}),
	"feed":       reflect.TypeOf(feed{}),
	"first_name": reflect.TypeOf(firstName{}),
	"friends":    reflect.TypeOf(friends{}),
}

type defaultStruct struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type about struct {
	defaultStruct
}

type activities struct {
	Field string `json:"field"`
	// TODO : make it as enum {add, change, edit, edited, delete, follow, hide, receive, remove, send, share, unhide, update}
	Verb  string `json:"verb"`
	Value struct {
		Page string `json:"page"`
	} `json:"value"`
}

type birthday struct {
	defaultStruct
}

type books struct {
	Field string `json:"field"`
	// TODO : make it as enum {add, change, edit, edited, delete, follow, hide, receive, remove, send, share, unhide, update}
	Verb  string `json:"verb"`
	Value struct {
		Page string `json:"page"`
	} `json:"value"`
}

type education struct {
	Type      string                `json:"type"`
	Education []educationExperience `json:"value"`
}

type email struct {
	defaultStruct
}

type events struct {
	Field string `json:"field"`
	Value struct {
		EventID string `json:"event_id"`
		// TODO : make it as enum {accept, create, decline, maybe, update}
		Verb string `json:"verb"`
	} `json:"value"`
}

type feed struct {
	Field string `json:"field"`
}

type firstName struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type friends struct {
	Field string `json:"field"`
	Value struct {
		UUID string `json:"uuid"`
	} `json:"value"`
	//TODO : make it as enum {add, change, edit, edited, delete, follow, hide, receive, remove, send, share, unhide, update}
	Verb string `json:"verb"`
}

type educationExperience struct {
	ID            string `json:"id"`
	Classes       []Experience
	Concentration []Page
	Degree        Page
	School        Page
	With          []User
	Year          Page
	Field         string `json:"field"`
}

type Experience struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	From        User   `json:"from"`
	Name        string `json:"name"`
	With        []User `json:"with"`
}

type User struct {
}
