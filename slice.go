http://www.golangbootcamp.com/book/collection_types

Note however, that you would get a runtime error if you were to do that:

cities := []string{}
cities[0] = "Santa Monica"
As explained above, a slice is seating on top of an array, in this case, the array is empty and the slice can’t set a value in the referred array. There is a way to do that though, and that is by using the append function:

package main

import "fmt"

func main() {
	cities := []string{}
	cities = append(cities, "San Diego")
	fmt.Println(cities)
	// [San Diego]
