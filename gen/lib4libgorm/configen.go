package lib4libgorm

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportConfigForLibgorm ...
func ExportConfig(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
