package time

import (
	"errors"
	"hbcase/aggregate"
)

var (
	ErrorTimeNotFound     = errors.New("time not found")
	ErrFailedToAddTime    = errors.New("failed to add the time to the repository")
	ErrFailedToUpdateTime = errors.New("failed to update the time to the repository")
)

type TimeRepository interface {
	Get() (aggregate.Time, error)
	Save(t aggregate.Time) error
	IncreaseTime(t aggregate.Time) (aggregate.Time, error)
}
