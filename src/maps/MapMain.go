package main

import (
	"fmt"
)

var family map[string]string = map[string]string{"name" : "bizhenchao"}

func setName(obj map[string]string) {
	obj["name"] = "222"
}

func main() {
	setName(family)
	fmt.Printf("hte name is %s", family["name"])
}

