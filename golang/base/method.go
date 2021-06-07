package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

// receiver: m *Movie
// func name: summary
// return type: string
func (m *Movie) toString() string {
	ret := fmt.Sprintf("%s:%s,%s:%.1f",
		"Name", m.Name,
		"Rating", m.Rating)
	return ret
}

// m is value refrence
func (m Movie) changeRating(rating float32) {
	fmt.Println("changeRating to ", rating)
	m.Rating = rating
}

func main() {
	m := Movie{
		Name:   "Go",
		Rating: 8.8,
	}
	fmt.Println(m.toString())

	m.changeRating(7.7)
	fmt.Println(m) // 8.8
}
