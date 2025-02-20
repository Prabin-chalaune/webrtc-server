package dtos

type SignalingMessage struct {
	Type    string      `json:"type"`
	Room    string      `json:"room"`
	Payload interface{} `json:"payload"`
}
