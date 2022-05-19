package models

type Login struct {
	Model
	Client string `json:"client"`
	IP     string `json:"ip"`
	Type   string `json:"type"`
}
