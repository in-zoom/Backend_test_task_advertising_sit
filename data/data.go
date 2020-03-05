package data

type NewAd struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Price       string `json:"price"`
}

type Ads struct {
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Price       string   `json:"price"`
	Data        string   `json:"data"`
	Link        []string `json: "link"`
}
