package types

type Cursor struct {
	After  *string `json:"after"`
	Before *string `json:"before"`
	Limit  *int `json:"limit"`
}
