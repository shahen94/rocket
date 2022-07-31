package rocket

type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	sessionType string
}

type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}
