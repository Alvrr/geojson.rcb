package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"https://example.com",
	"http://localhost:5174/",
	"http://localhost:5173/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
