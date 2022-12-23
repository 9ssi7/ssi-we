package entity

import "time"

type Event struct {
	UUID      string      `json:"uuid,omitempty" bson:"_id,omitempty"`
	Page      string      `json:"page,omitempty" bson:"page"`
	Element   string      `json:"element,omitempty" bson:"element"`
	Type      string      `json:"type,omitempty" bson:"type"`
	Name      string      `json:"name,omitempty" bson:"name"`
	Params    interface{} `json:"params,omitempty" bson:"params,omitempty"`
	User      interface{} `json:"user,omitempty" bson:"user,omitempty"`
	IP        string      `json:"ip,omitempty" bson:"ip"`
	CreatedAt time.Time   `json:"dateOfCreate,omitempty" bson:"created_at"`
}

func (e *Event) SetUUID(uuid string) {
	e.UUID = uuid
}
