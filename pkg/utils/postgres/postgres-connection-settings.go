package postgres

type ConnectionSettings struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}
