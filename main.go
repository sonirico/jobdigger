package main

import (
	"fmt"
	"jobdigger/digger/empleoPonferradaOrg"
)

func main () {
	empleoPonferradaOrgDigger := empleoPonferradaOrg.New("https://empleo.ponferrada.org/rss")
	offers := empleoPonferradaOrgDigger.FetchAll()

	for index, offer := range offers {
		fmt.Println(index, offer)
	}
}
