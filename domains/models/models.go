package models

type Config struct {
	Application struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}
	Database struct {
		Type         string
		Host         string
		Port         int
		User         string `config:"envVar"`
		Password     string `config:"envVar"`
		Dbname       string `config:"envVar"`
		MaxIdleConns int
		MaxOpenConns int
	}
}
