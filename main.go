package main

import (
	"fmt"
	"jobdigger/digger/empleoPonferradaOrg"
)

func main() {
	empleoPonferradaOrgDigger := empleoPonferradaOrg.New("https://empleo.ponferrada.org/rss")
	result, err := empleoPonferradaOrgDigger.Get()

	if err != nil {
		fmt.Println("Fetch error")
	}

	for index, offer := range result.Offers {
		fmt.Println(index, offer)
	}

	fmt.Println(fmt.Sprintf("OK: %d \nFAILED: %d\nTOTAL: %d", result.Ok, result.Failed, result.Total()))
}
