package types

type Providers struct {
	Providers []Provider `json:"providers"`
}

type Provider struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
