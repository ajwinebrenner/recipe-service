package config

type Config struct {
	DatabaseURL string
	Host        string
	Port        string
}

func Load() (*Config, error) {
	// could use a package to load defaults in
	// no real need for error at this point
	c := Config{
		DatabaseURL: "postgresql://user:password@localhost:5432/recipe?sslmode=disable",
		Host:        "0.0.0.0",
		Port:        "8080",
	}
	return &c, nil
}
