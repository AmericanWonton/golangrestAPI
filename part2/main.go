package main

import (
	"fmt"
)

func main() {
	coralSlice := []string{"blue coral", "foliose coral", "pillar coral", "elkhorn coral", "black coral", "antipathes", "leptopsammia", "massive coral", "soft coral"}

	coralSlice = append(coralSlice[:3], coralSlice[4:]...)

	fmt.Printf("%q\n", coralSlice)
}
