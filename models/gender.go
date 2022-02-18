package models

import "fmt"

// Gender is Enum used in users and drivers gender
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (e *Gender) Scan(gender interface{}) error {
	switch s := gender.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", gender)
	}
	return nil
}
