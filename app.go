package brass

type App struct {
	Icon     []byte // 1024x1024 pixel .png file
	Name     string
	CoreType string
	Types    map[string]Type
}
