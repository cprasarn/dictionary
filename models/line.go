package models

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type Source struct {
	Type    string `json:"type"`
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

type Message struct {
	ID   *string `json:"id"`
	Type string  `json:"type"`
	Text string  `json:"text"`
}
type Event struct {
	Type            string           `json:"type"`
	Message         *Message         `json:"message"`
	ReplyToken      string           `json:"replyToken"`
	Source          *Source          `json:"source"`
	Mode            *string          `json:"mode"`
	DeliveryContext *DeliveryContext `json:"deliveryContext"`
}

type Input struct {
	Destination *string `json:"destination"`
	Events      []Event `json:"events"`
}
