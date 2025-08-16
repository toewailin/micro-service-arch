package config

// Config structure to hold configuration values
type Config struct {
	DatabaseFile string
}

// LoadConfig loads configuration from environment variables or defaults
func LoadConfig() *Config {
	// In a real application, you would load from environment variables or config files
	return &Config{
		DatabaseFile: "product.db", // SQLite file for this example
	}
}
