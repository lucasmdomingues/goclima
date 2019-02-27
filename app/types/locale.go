package types

type Locale struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Country string `json:"country"`
}
