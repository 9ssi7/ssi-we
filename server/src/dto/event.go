package dto

type EventCreateRequest struct {
	Page    string      `json:"page,omitempty" bson:"page"`
	Element string      `json:"element,omitempty" bson:"element"`
	Type    string      `json:"type,omitempty" bson:"type"`
	Name    string      `json:"name,omitempty" bson:"name"`
	Params  interface{} `json:"params,omitempty" bson:"params,omitempty"`
	User    interface{} `json:"user,omitempty" bson:"user,omitempty"`
	IP      string
}
