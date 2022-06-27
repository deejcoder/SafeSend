package interfaces

type EntityType string

const (
	User  EntityType = "USER"
	Group EntityType = "GROUP"
)

// Values returns a list of valid Enum values
func (EntityType) Values() (kinds []string) {
	for _, s := range []EntityType{User, Group} {
		kinds = append(kinds, string(s))
	}
	return
}
