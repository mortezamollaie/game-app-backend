package entity

type Category string

const (
	FootballCategory   = "Football"
	TechnologyCategory = "Technology"
)

func (c Category) IsValid() bool {
	switch c {
	case FootballCategory, TechnologyCategory:
		return true
	}
	return false
}
