package models

type CreateResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		ID string `json:"id"`
	} `json:"data"`
	CustomMessage string `json:"custom_message"`
}

type Table struct {
	ID     string            `json:"id"`
	Name   string            `json:"name"`
	Fields map[string]string `json:"fields"`
}

type AllTable struct {
	Tables []*Table `json:"tables"`
}

type FieldResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Fields []map[string]interface {
		} `json:"fields"`
		Count int `json:"count"`
	}
}
