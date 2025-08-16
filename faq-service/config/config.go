package config

// Config structure to hold configuration values
type Config struct {
	DatabaseFile string
}

// LoadConfig loads configuration from environment variables or defaults
func LoadConfig() *Config {
	// In a real application, you could load from environment variables or a config file
	return &Config{
		DatabaseFile: "faq.db", // SQLite file for this example
	}
}
