package webhook

import (
	"bytes"
	"encoding/json"
	"reflect"

	"golang-fb-api/core"
	"strings"
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
			if a.Field == "pic" || strings.Contains(a.Field, "pic_") || a.Field == "picture" {
				a.Field = "pic"
			}
			changes = append(changes, TypedChange{Type: core.UserTypesMap[a.Field], Data: rawValue})
		}
	}
	return changes
}
