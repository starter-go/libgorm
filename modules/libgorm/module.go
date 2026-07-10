package libgorm

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgorm"
	"github.com/starter-go/libgorm/gen/demo4libgorm"
	"github.com/starter-go/libgorm/gen/lib4libgorm"
	"github.com/starter-go/libgorm/gen/test4libgorm"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/modules/units"
)

// Module 导出模块 ['github.com/starter-go/libgorm']
func Module() application.Module {
	mb := libgorm.SrcLibModuleBuilder()
	mb.Components(lib4libgorm.ExportConfig)

	mb.Depend(starter.Module())

	return mb.Create()
}

// Module 导出模块 ['github.com/starter-go/libgorm']
func ModuleForTest() application.Module {
	mb := libgorm.SrcTestModuleBuilder()
	mb.Components(test4libgorm.ExportConfig)

	mb.Depend(Module())
	mb.Depend(units.Module())

	return mb.Create()
}

// Module 导出模块 ['github.com/starter-go/libgorm']
func ModuleForDemo() application.Module {
	mb := libgorm.SrcDemoModuleBuilder()
	mb.Components(demo4libgorm.ExportConfig)

	mb.Depend(Module())

	return mb.Create()
}
