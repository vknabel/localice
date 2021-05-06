package internal

type Localization struct {
	Name         string
	Translations []Translation
}
type Translation struct {
	Platform string
	Key      string
	Text     string
}
