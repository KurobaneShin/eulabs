package types

type Product struct {
	Id          int    `bun:"id,pk,autoincrement" json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
}
