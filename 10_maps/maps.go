package main

import (
	"fmt"
	"maps"
)

// maps -> hashmap in cpp, objects in js, dictionary in python
func main() {
	// creating map
	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])
	// fmt.Println(m["phone"]) // if key does not exists in the map it returns zero value

	// m := make(map[string]int)
	// m["age"] = 21
	// fmt.Println(m)
	// fmt.Println(m["age"])
	// fmt.Println(m["phone"]) // prints 0 becuse key does not exists

	// m := make(map[string]int)
	// m["age"] = 21
	// m["price"] = 50
	// fmt.Println(m)
	// fmt.Println(len(m))

	// delete(m, "price")
	// fmt.Println(m)

	// clear(m)
	// fmt.Println(m)

	// m := map[string]int{"price":40,"phones":3}
	// fmt.Println(m)
	// v, ok := m["phones"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not okay")
	// }

	m1 := map[string]int{"price":40, "phones":3}
	m2 := map[string]int{"price":40, "phones":8}
	fmt.Println(maps.Equal(m1, m2))
}