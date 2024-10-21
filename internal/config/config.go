package config

import (
	"flag"
	"os"
)

type Config struct {
	RunAddress  string
	BrokersList string
	//DatabaseURI          string
	//DatabaseName         string
	//AccrualSystemAddress string
}

func Load() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.RunAddress, "a", ":8080", "Address and port to run server")
	flag.StringVar(&cfg.BrokersList, "b", "kafka:9092", "Kafka connection URI")
	//flag.StringVar(&cfg.DatabaseURI, "d", "mongodb://localhost:27017", "MongoDB connection URI")
	//flag.StringVar(&cfg.DatabaseName, "db", "gophermartLoyalty", "MongoDB database name")
	//flag.StringVar(&cfg.AccrualSystemAddress, "r", "", "Accrual system address")
	flag.Parse()

	if envRunAddress := os.Getenv("RUN_ADDRESS"); envRunAddress != "" {
		cfg.RunAddress = envRunAddress
	}
	if envBrokersList := os.Getenv("kafka:9092"); envBrokersList != "" {
		cfg.BrokersList = envBrokersList
	}
	//if envDatabaseURI := os.Getenv("DATABASE_URI"); envDatabaseURI != "" {
	//	cfg.DatabaseURI = envDatabaseURI
	//}
	//if envDatabaseName := os.Getenv("DATABASE_NAME"); envDatabaseName != "" {
	//	cfg.DatabaseName = envDatabaseName
	//}
	//if envAccrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS"); envAccrualSystemAddress != "" {
	//	cfg.AccrualSystemAddress = envAccrualSystemAddress
	//}

	return cfg, nil
}
