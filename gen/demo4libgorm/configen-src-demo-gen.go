package demo4libgorm
import (
    p512a30914 "github.com/starter-go/libgorm"
    p62a9d7674 "github.com/starter-go/libgorm/src/demo/golang/demos/demo1"
     "github.com/starter-go/application"
)

// type p62a9d7674.TaDaoImpl in package:github.com/starter-go/libgorm/src/demo/golang/demos/demo1
//
// id:com-62a9d76740f72c9d-demo1-TaDaoImpl
// class:
// alias:
// scope:singleton
//
type p62a9d76740_demo1_TaDaoImpl struct {
}

func (inst* p62a9d76740_demo1_TaDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-62a9d76740f72c9d-demo1-TaDaoImpl"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p62a9d76740_demo1_TaDaoImpl) new() any {
    return &p62a9d7674.TaDaoImpl{}
}

func (inst* p62a9d76740_demo1_TaDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p62a9d7674.TaDaoImpl)
	nop(ie, com)

	
    com.Src = inst.getSrc(ie)


    return nil
}


func (inst*p62a9d76740_demo1_TaDaoImpl) getSrc(ie application.InjectionExt)p62a9d7674.MyAgent{
    return ie.GetComponent("#alias-62a9d76740f72c9d18d20b8bc1ba693a-MyAgent").(p62a9d7674.MyAgent)
}



// type p62a9d7674.TbDaoImpl in package:github.com/starter-go/libgorm/src/demo/golang/demos/demo1
//
// id:com-62a9d76740f72c9d-demo1-TbDaoImpl
// class:
// alias:
// scope:singleton
//
type p62a9d76740_demo1_TbDaoImpl struct {
}

func (inst* p62a9d76740_demo1_TbDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-62a9d76740f72c9d-demo1-TbDaoImpl"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p62a9d76740_demo1_TbDaoImpl) new() any {
    return &p62a9d7674.TbDaoImpl{}
}

func (inst* p62a9d76740_demo1_TbDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p62a9d7674.TbDaoImpl)
	nop(ie, com)

	
    com.Src = inst.getSrc(ie)


    return nil
}


func (inst*p62a9d76740_demo1_TbDaoImpl) getSrc(ie application.InjectionExt)p62a9d7674.MyAgent{
    return ie.GetComponent("#alias-62a9d76740f72c9d18d20b8bc1ba693a-MyAgent").(p62a9d7674.MyAgent)
}



// type p62a9d7674.TcDaoImpl in package:github.com/starter-go/libgorm/src/demo/golang/demos/demo1
//
// id:com-62a9d76740f72c9d-demo1-TcDaoImpl
// class:
// alias:
// scope:singleton
//
type p62a9d76740_demo1_TcDaoImpl struct {
}

func (inst* p62a9d76740_demo1_TcDaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-62a9d76740f72c9d-demo1-TcDaoImpl"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p62a9d76740_demo1_TcDaoImpl) new() any {
    return &p62a9d7674.TcDaoImpl{}
}

func (inst* p62a9d76740_demo1_TcDaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p62a9d7674.TcDaoImpl)
	nop(ie, com)

	
    com.Src = inst.getSrc(ie)


    return nil
}


func (inst*p62a9d76740_demo1_TcDaoImpl) getSrc(ie application.InjectionExt)p62a9d7674.MyAgent{
    return ie.GetComponent("#alias-62a9d76740f72c9d18d20b8bc1ba693a-MyAgent").(p62a9d7674.MyAgent)
}



// type p62a9d7674.TableReg in package:github.com/starter-go/libgorm/src/demo/golang/demos/demo1
//
// id:com-62a9d76740f72c9d-demo1-TableReg
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-62a9d76740f72c9d18d20b8bc1ba693a-MyAgent
// scope:singleton
//
type p62a9d76740_demo1_TableReg struct {
}

func (inst* p62a9d76740_demo1_TableReg) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-62a9d76740f72c9d-demo1-TableReg"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-62a9d76740f72c9d18d20b8bc1ba693a-MyAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p62a9d76740_demo1_TableReg) new() any {
    return &p62a9d7674.TableReg{}
}

func (inst* p62a9d76740_demo1_TableReg) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p62a9d7674.TableReg)
	nop(ie, com)

	
    com.DSMan = inst.getDSMan(ie)
    com.Prefix = inst.getPrefix(ie)
    com.URI = inst.getURI(ie)
    com.SourceName = inst.getSourceName(ie)


    return nil
}


func (inst*p62a9d76740_demo1_TableReg) getDSMan(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}


func (inst*p62a9d76740_demo1_TableReg) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("datagroup.default.table-name-prefix")
}


func (inst*p62a9d76740_demo1_TableReg) getURI(ie application.InjectionExt)string{
    return ie.GetString("datagroup.default.uri")
}


func (inst*p62a9d76740_demo1_TableReg) getSourceName(ie application.InjectionExt)string{
    return ie.GetString("datagroup.default.datasource")
}


