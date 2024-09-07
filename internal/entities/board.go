package entities

type Board struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Lists []List `json:"lists"`
}
