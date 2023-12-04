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
	Name   string            `json:"name"`
	Fields map[string]string `json:"fields"`
}

type AllTable struct {
	Tables []*Table `json:"tables"`
}
