package main

import (
	"os"

	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/starter"
)

func main() {

	a := os.Args
	m := libgorm.ModuleForTest()
	i := starter.Init(a)

	i.MainModule(m)
	i.WithPanic(true).Run()
}
