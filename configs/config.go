package configs

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
	RabbitMQUser      string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword  string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost      string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort      string `mapstructure:"RABBITMQ_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}

func MakeMigrations(path string, db *sql.DB) error {
	fmt.Println("Making migrations...")

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		// Duas barras path relativo
		// Três barras path absoluto
		fmt.Sprintf("file://%s", path),
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		switch err.Error() {
		case "no change":
			fmt.Println("No migrations changes.")
			return nil
		default:
			return err
		}
	}

	fmt.Println("Migrations finished.")
	return nil
}
