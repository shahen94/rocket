package config

type Config struct {
	Port        string
	Renderer    string
	Cookie      CookieConfig
	SessionType string
}
