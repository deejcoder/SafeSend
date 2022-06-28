package interfaces

type TokenProvider string

const (
	Google TokenProvider = "GOOGLE"
)

// Values returns a list of valid Enum values
func (TokenProvider) Values() (kinds []string) {
	for _, s := range []TokenProvider{Google} {
		kinds = append(kinds, string(s))
	}
	return
}
