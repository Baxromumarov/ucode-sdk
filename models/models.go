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
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Slug   string  `json:"slug"`
	Fields []Field `json:"field_type"` // [field_name] = field_type
}
type Field struct {
	Name string
	Slug string
	Type string
}
type FieldResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Fields []map[string]interface {
		} `json:"fields"`
	}
}

type CreateMenuResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Label      string `json:"label"`
		Icon       string `json:"icon"`
		ParentID   string `json:"parent_id"`
		Type       string `json:"type"`
		ID         string `json:"id"`
		Order      int    `json:"order"`
		Attributes struct {
			Label   string `json:"label"`
			LabelEn string `json:"label_en"`
		} `json:"attributes"`
		Data struct {
			Permission struct {
				Delete       bool   `json:"delete"`
				GUID         string `json:"guid"`
				MenuID       string `json:"menu_id"`
				MenuSettings bool   `json:"menu_settings"`
				Read         bool   `json:"read"`
				RoleID       string `json:"role_id"`
				Update       bool   `json:"update"`
				Write        bool   `json:"write"`
			} `json:"permission"`
			Table struct {
				Attributes         any    `json:"attributes"`
				Description        string `json:"description"`
				FolderID           string `json:"folder_id"`
				Icon               string `json:"icon"`
				ID                 string `json:"id"`
				IsCached           bool   `json:"is_cached"`
				IsChanged          bool   `json:"is_changed"`
				IsChangedByHost    any    `json:"is_changed_by_host"`
				IsLoginTable       bool   `json:"is_login_table"`
				IsSystem           bool   `json:"is_system"`
				Label              string `json:"label"`
				OrderBy            bool   `json:"order_by"`
				SectionColumnCount int    `json:"section_column_count"`
				ShowInMenu         bool   `json:"show_in_menu"`
				Slug               string `json:"slug"`
				SoftDelete         bool   `json:"soft_delete"`
				SubtitleFieldSlug  string `json:"subtitle_field_slug"`
				WithIncrementID    bool   `json:"with_increment_id"`
			} `json:"table"`
		} `json:"data"`
	} `json:"data"`
	CustomMessage string `json:"custom_message"`
}

type RelationResponse struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Data        struct {
		Type       string `json:"type"`
		Attributes struct {
			Format    string `json:"format"`
			Label     string `json:"label"`
			LabelEn   string `json:"label_en"`
			LabelTo   string `json:"label_to"`
			LabelToEn string `json:"label_to_en"`
			Math      struct {
				Label string `json:"label"`
				Value string `json:"value"`
			} `json:"math"`
			Options []any `json:"options"`
		} `json:"attributes"`
		ID string `json:"id"`
	} `json:"data"`
	CustomMessage string `json:"custom_message"`
}
