package brass

type App struct {
	Repo     string
	Name     string
	Icon     []byte // 1024x1024 pixel .png file
	CoreType string
	Types    map[string]Type
}
