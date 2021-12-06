package aggregate

import "hbcase/entity"

type Time struct {
	item *entity.Time
}

func NewTime(time int) Time {
	return Time{item: &entity.Time{Time: time, Id: 1}}
}

func (t *Time) GetTime() int {
	if t.item == nil {
		return 1
	}
	return t.item.Time
}

func (t *Time) GetId() int {
	return t.item.Id
}

func (t *Time) SetTime(time int) {
	t.item.Time = time
}
