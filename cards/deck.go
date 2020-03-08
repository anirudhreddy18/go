package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, car := range d {
		fmt.Println(i, car)
	}
}
