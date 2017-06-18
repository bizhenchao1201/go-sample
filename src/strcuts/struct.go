package structs

import (
	"fmt"
)

type Family struct {
	father string
	faAge int
	mother string
	moAge int	
}s

//var bizhenchao = Family{"bizhdenchao", 35, "sunhuimin", 30}

func setAge(family Family) {
	family.faAge = family.faAge + 1
	family.moAge = family.moAge + 1
}

func main() {
	bizhenchao := Family{"bizhdenchao", 35, "sunhuimin", 30}
	
	setAge(bizhenchao)
	
	fmt.Printf("father age is %d \n", bizhenchao.faAge)
	fmt.Printf("mathor age is %d", bizhenchao.moAge)
	
}



