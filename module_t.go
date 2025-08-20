package libgorm

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/starter-go/libgorm"
	theModuleVersion  = "v0.9.14"
	theModuleRevision = 15
)

////////////////////////////////////////////////////////////////////////////////

const (
	theSrcLibModuleResPath  = "src/lib/resources"
	theSrcDemoModuleResPath = "src/demo/resources"
	theSrcTestModuleResPath = "src/test/resources"
)

//go:embed "src/lib/resources"
var theSrcLibModuleResFS embed.FS

//go:embed "src/demo/resources"
var theSrcDemoModuleResFS embed.FS

//go:embed "src/test/resources"
var theSrcTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// ModuleBuilder 用于创建模块 ['github.com/starter-go/libgorm']
func SrcLibModuleBuilder() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#lib")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theSrcLibModuleResFS, theSrcLibModuleResPath)

	return mb
}

// ModuleBuilder 用于创建模块 ['github.com/starter-go/libgorm']
func SrcDemoModuleBuilder() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#demo")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theSrcDemoModuleResFS, theSrcDemoModuleResPath)

	return mb
}

// ModuleBuilder 用于创建模块 ['github.com/starter-go/libgorm']
func SrcTestModuleBuilder() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theSrcTestModuleResFS, theSrcTestModuleResPath)

	return mb
}
