package main

import "fmt"

func main() {
	countries := []string{"china", "usa", "india", "japan"}
	neededCountry := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountry))
	copy(countriesCpy, neededCountry)
	fmt.Println(countriesCpy)

}
