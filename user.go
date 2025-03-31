package brass

type User struct {
	ID           string
	PasswordHash string
	Sessions     map[string]bool
}
