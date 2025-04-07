package entities

type Measurements struct {
	ID            int     `json:"id"`
	Id_device     string  `json:"id_device"`
	Temp          float32 `json:"temp"`
	Humedity      float32 `json:"humedity"`
	Air           float32 `json:"air"`
	Sun           float32 `json:"sun"`
	Date_and_hour string  `json:date_and_hour`
}
