package model

// import (
// 	"fmt"
// 	"strconv"
// )

type Movie struct {
	Id          string `json:id`
	Title       string `json:"title"`
	Director    string `json:director`
	ReleaseDate int    `json:releaseDate`
}

// func (m *Movie) toMap() map[string]interface{} {
// 	toMap := map[string]interface{}{}
// 	toMap["title"] = m.Title
// 	toMap["director"] = m.Director
// 	toMap["releaseDate"] = m.ReleaseDate
// 	return toMap
// }

// func FromMap(json map[string]interface{}) *Movie {
// 	title := json["title"]
// 	director := json["director"]
// 	releaseDate, _ := strconv.Atoi(fmt.Sprint(json["releseDate"]))
// 	return &Movie{fmt.Sprint(title), fmt.Sprint(director), releaseDate}
// }
