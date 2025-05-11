package model

type EventBridgeDetail struct {
	Object struct {
		Key string `json:"key"`
	} `json:"object"`
}

type EventBridgeEvent struct {
	Detail EventBridgeDetail `json:"detail"`
}
