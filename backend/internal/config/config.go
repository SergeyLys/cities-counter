package config

import "os"

type Config struct {
	Port             string
	GeoNamesURL      string
	GeoNamesUsername string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		Port:             port,
		GeoNamesURL:      "http://api.geonames.org/searchJSON",
		GeoNamesUsername: "hsample",
	}
}
