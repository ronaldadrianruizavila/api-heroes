package heroes

type Response struct {
	Data Heroes `json:"data"`
}

type Heroes struct {
	Results []Heroe `json:"results"`
}

type Heroe struct {
	Id          int    `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
}
