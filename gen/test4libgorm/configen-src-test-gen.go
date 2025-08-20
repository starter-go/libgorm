package test4libgorm
import (
    pac62b8721 "github.com/starter-go/libgorm/src/test/golang/units4libgorm"
     "github.com/starter-go/application"
)

// type pac62b8721.Unit1 in package:github.com/starter-go/libgorm/src/test/golang/units4libgorm
//
// id:com-ac62b8721b25bb3e-units4libgorm-Unit1
// class:
// alias:
// scope:singleton
//
type pac62b8721b_units4libgorm_Unit1 struct {
}

func (inst* pac62b8721b_units4libgorm_Unit1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ac62b8721b25bb3e-units4libgorm-Unit1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pac62b8721b_units4libgorm_Unit1) new() any {
    return &pac62b8721.Unit1{}
}

func (inst* pac62b8721b_units4libgorm_Unit1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pac62b8721.Unit1)
	nop(ie, com)

	


    return nil
}


