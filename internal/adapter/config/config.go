package config

import "os"

// Container contains environment variables for the application, database, cache, token, and http server
type (
	Container struct {
		App  *App
		DB   *DB
		HTTP *HTTP
	}
	// App contains all the environment variables for the application
	App struct {
		Name string
		Env  string
	}

	// Database contains all the environment variables for the database
	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
	// HTTP contains all the environment variables for the http server
	HTTP struct {
		Env  string
		URL  string
		Port string
	}
)

// New creates a new container instance
func New() (*Container, error) {

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Env:  os.Getenv("APP_ENV"),
		URL:  os.Getenv("HTTP_URL"),
		Port: os.Getenv("HTTP_PORT"),
	}

	return &Container{
		app,
		db,
		http,
	}, nil
}
