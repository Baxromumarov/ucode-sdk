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
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	FieldType map[string]string `json:"field_type"` // [field_name] = field_type
	FieldID   map[string]string `json:"field_id"`   // [field_name] = field_id
}

type AllTable struct {
	Tables []Table `json:"tables"`
}
