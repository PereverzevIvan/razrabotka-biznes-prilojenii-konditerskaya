package models

type ToolWorkInterval struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Description string `json:"description"`
}
