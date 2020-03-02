package data

type NewAd struct {
    Id          int      `json:"id"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Price       string   `json:"price"`
	Link        []string `json: "link"`
}
