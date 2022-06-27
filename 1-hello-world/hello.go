package helloworld

var Prefixes = map[string]string{
	"english": "Hello",
	"spanish": "Hola",
	"french":  "Bonjour",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	prefix, exists := Prefixes[language]
	if exists {
		return prefix + ", "
	}

	return "Hello, "
}
