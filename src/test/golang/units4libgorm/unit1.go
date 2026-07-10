package units4libgorm

import (
	"context"

	"github.com/starter-go/libgorm"
	"github.com/starter-go/units"
)

type Unit1 struct {

	//starter:component

	DSM libgorm.DataSourceManager //starter:inject("#")
}

// ListRegistrations implements units.Unit.
func (inst *Unit1) ListRegistrations(list []*units.Registration) []*units.Registration {

	u1 := &units.Registration{
		Name:    "unit1",
		Enabled: true,
		Do:      inst.run,
	}

	list = append(list, u1)

	return list
}

func (inst *Unit1) run(cc context.Context) error {

	const (
		alias = "demo1ds"
	)

	ds, err := inst.DSM.GetDataSource(alias)
	if err != nil {
		return err
	}

	db, err := ds.DB()
	if err != nil {
		return err
	}

	var count int64
	db.Count(&count)

	return nil
}

func (inst *Unit1) _impl() units.Unit {
	return inst
}
