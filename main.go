package main

import (
	"fmt"
	worker2 "jobdigger/digger"
)

func main () {
	worker := worker2.New("https://empleo.ponferrada.org/rss")
	offers := worker.FetchAll()
	for index, offer := range offers {
		fmt.Println(index, offer)
	}
}
