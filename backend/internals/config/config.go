package config

type configStruct struct {
	ServerPort int
	DBUri      string
}

var config *configStruct = nil

func GetConfig() *configStruct {
	if config == nil {
		config = &configStruct{
			ServerPort: 8080,
			// DBUri:      "postgres://admin:secret@movie_db:5432/movie_db ",
			DBUri: "postgres://admin:secret@localhost:5432/movie_db?sslmode=disable",
		}
	}
	return config
}
