package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"golang-fb-api/core"
)

type Change struct {
	json.RawMessage
}

// TypedChange struct for storing dynamic data
// with raw value and it`s type
type TypedChange struct {
	Type reflect.Type
	Data []byte
}

// Entry struct contains incoming data from facebook
type Entry struct {
	Time    int64             `json:"time"`
	Changes []json.RawMessage `json:"changes"`
	ID      string            `json:"id"`
	UID     string            `json:"uid"`
}

// Event facebook wrapper around entry
type Event struct {
	Entry  []Entry `json:"entry"`
	Object string  `json:"object"`
}

func (e *Entry) GetChanges() []TypedChange {
	changes := make([]TypedChange, 0)
	var a = struct {
		Field string `json:"field"`
	}{}
	for _, v := range e.Changes {
		rawValue, _ := v.MarshalJSON()
		if string(rawValue) != "null" {
			json.NewDecoder(bytes.NewBuffer(rawValue)).Decode(&a)
			fmt.Println(a.Field)
			changes = append(changes, TypedChange{Type: core.UserTypesMap[a.Field], Data: rawValue})
		}
	}
	return changes
}
