package internaldb

import (
	"hbcase/aggregate"
	"hbcase/domain/time"

	db "github.com/sonyarouje/simdb"
)

type ScribbleRepository struct {
	db *db.Driver
}

type simdbTime struct {
	Id   int
	Time int
}

func (c simdbTime) ID() (jsonField string, value interface{}) {
	value = c.Id
	jsonField = "Id"
	return
}

func NewFromTime(c aggregate.Time) simdbTime {
	return simdbTime{
		Time: c.GetTime(),
		Id:   c.GetId(),
	}
}

func New() (*ScribbleRepository, error) {
	driver, err := db.New("ecomm")
	if err != nil {
		return nil, err
	}
	db := driver.Open(simdbTime{})
	return &ScribbleRepository{db: db}, nil
}

func (m simdbTime) ToAggregate() aggregate.Time {
	c := aggregate.NewTime(m.Time)
	return c
}

func (mr *ScribbleRepository) Get() (aggregate.Time, error) {
	p, e := mr.first()
	if e != nil {
		return aggregate.Time{}, time.ErrorTimeNotFound
	}
	return p.ToAggregate(), nil
}

func (mr *ScribbleRepository) first() (simdbTime, error) {
	p := simdbTime{}
	e := mr.db.First().AsEntity(&p)
	if e != nil {
		return simdbTime{}, time.ErrorTimeNotFound
	}
	return p, nil
}

func (mr *ScribbleRepository) Save(t aggregate.Time) error {
	p := NewFromTime(t)
	println(p.Id)
	e := mr.db.Insert(p)
	if e != nil {
		return time.ErrFailedToAddTime
	}
	return nil
}

func (mr *ScribbleRepository) IncreaseTime(model aggregate.Time) (aggregate.Time, error) {

	tt, e := mr.first()
	if e != nil {
		se := mr.Save(model)
		if se != nil {
			return model, se
		}
	}
	tt.Time = tt.Time + model.GetTime()

	err := mr.db.Update(tt)
	if err != nil {
		println(err.Error())
		return aggregate.Time{}, time.ErrFailedToUpdateTime
	}
	return tt.ToAggregate(), nil

}
