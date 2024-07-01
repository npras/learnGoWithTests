package hello

const (
	tamil = "Tamil"
	hindi = "Hindi"

	engHelloPrefix   = "Hello "
	tamilHelloPrefix = "Vaadaa "
	hindiHelloPrefix = "Aayiye "
)

// This guy says Hello.
func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

// This guys greets in many lanbgs.
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case tamil:
		prefix = tamilHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}
