package settings

type SettingsProvider interface {
	Get(key string) (string, error)
	Set(key, value string) error
}
