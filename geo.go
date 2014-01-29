package geotrigger

import "encoding/json"


type Geo interface {
	MarshalJSON() ([]byte, error)
}

type Point struct {
	Lat      float64 `json:"latitude"`
	Lng      float64 `json:"longitude"`
	Distance float64 `json:"distance"`
}

func (p Point) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		geo Point `json:"geo"`
	}{p})
}
