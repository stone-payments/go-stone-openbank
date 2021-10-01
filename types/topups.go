package types

type Providers struct {
	Providers []Provider `json:"providers"`
}

type Provider struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type Product struct {
	Value int    `json:"value"`
	Name  string `json:"product_name,omitempty"`
}

type Products struct {
	Products []Product `json:"product"`
}
