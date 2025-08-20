package test4libgorm

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportConfigForDemo1 ...
func ExportConfig(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
