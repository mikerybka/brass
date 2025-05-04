package brass

func isBuiltinType(typeID string) bool {
	switch typeID {
	case "string", "int", "bool", "text", "image", "video", "audio":
		return true
	default:
		return false
	}
}
