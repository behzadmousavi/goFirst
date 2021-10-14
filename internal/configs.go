package internal

type database struct {
	Username string `default:"postgres"`
	Password string `default:"newsecret"`
	Host     string `default:"127.0.0.1"`
	Port     int    `default:"5432"`
}
