package sources

type Stockinfo struct{
	Data    map[string]interface{}  `json:"data"`
	Message string                  `json:"message"`
	Status  int                     `json:"status"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

type Input struct{
Symbol string       `json:"symbol"`
}

type Output struct{
	Message string                  `json:"message"`
	Status  int                     `json:"status"`
	Data    interface{}             `json:"data"`
}
