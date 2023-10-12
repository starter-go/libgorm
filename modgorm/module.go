package modgorm

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/libgorm/gen/gen4demo1"
	"github.com/starter-go/libgorm/gen/gen4libgorm"
)

// Module 导出模块 ['github.com/starter-go/libgorm']
func Module() application.Module {
	mb := libgorm.ModuleBuilder()
	mb.Components(gen4libgorm.ExportConfigForLibgorm)
	return mb.Create()
}

// ModuleForTest 导出模块 ['github.com/starter-go/libgorm#demo1']
func ModuleForTest() application.Module {
	mb := libgorm.ModuleBuilder()
	mb.Name("github.com/starter-go/libgorm#demo1")
	mb.Components(gen4demo1.ExportConfigForDemo1)
	mb.Depend(Module())
	return mb.Create()
}
