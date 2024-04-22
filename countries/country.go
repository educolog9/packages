package countries

type Country string

// Country represents a country code.
const (
	DominicanRepublic Country = "do"
	UnitedStates      Country = "us"
	Default           Country = "do"
)

func (l Country) String() string {
	return string(l)
}
