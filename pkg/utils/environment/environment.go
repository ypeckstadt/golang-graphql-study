package environment

// Environment contains the structure for all env settings
type Environment struct {
	Database Database
	Server   Server
}

// Database holds data necessary for database configuration
type Database struct {
	Host     string `env:"DATABASE_HOST"`
	Port     int    `env:"DATABASE_PORT"`
	Username string `env:"DATABASE_USERNAME"`
	Password string `env:"DATABASE_PASSWORD"`
	Name     string `env:"DATABASE_NAME"`
}

type Server struct {
	Port string `env:"SERVER_PORT"`
}
