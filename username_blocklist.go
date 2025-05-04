package brass

var UsernameBlocklist = map[string]bool{
	"admin":  true,
	"api":    true,
	"public": true,
	"join":   true,
	"login":  true,
	"logout": true,
}
