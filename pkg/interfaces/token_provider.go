package interfaces

type TokenProvider string

const (
	Google   TokenProvider = "GOOGLE"
	SafeSend TokenProvider = "SAFE_SEND"
)

// Values returns a list of valid Enum values
func (TokenProvider) Values() (kinds []string) {
	for _, s := range []TokenProvider{Google, SafeSend} {
		kinds = append(kinds, string(s))
	}
	return
}
