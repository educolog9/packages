package languagues

type Language string

// English represents the English language.
const (
	English Language = "en"
	Spanish Language = "es"
	Default Language = "es"
)

func (l Language) String() string {
	return string(l)
}
