package entity

type Campaign struct {
	Name, ProductCode      string
	Duration, Limit, Count int
	Status                 bool
}
