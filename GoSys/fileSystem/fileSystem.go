package fileSystem

import (
	"fmt"
	_ "fmt"
	"github.com/ShiinaOrez/gowork/GoSys/typedefs"
)

func init() {
	fmt.Printf("You're now at [" + typedefs.ROOT + "], enjoy!\n")
}

func OutputBasePath() {
	fmt.Printf(typedefs.ROOT)
}
