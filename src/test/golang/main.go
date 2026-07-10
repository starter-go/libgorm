package main

import (
	"os"

	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/units"
)

func main() {

	a := os.Args
	m := libgorm.ModuleForTest()

	c := &units.Context{
		Arguments: a,
		Module:    m,
		UsePanic:  true,
	}

	units.Run(c)
}
