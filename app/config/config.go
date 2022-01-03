package config

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

const (
	configPath = "config/config.local.yaml"
)

type Configs struct {
	fx.Out
	DBConfig *DBConfig `yaml:"postgres"`

	DB *sqlx.DB
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	DBName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func SetupConfigs() (Configs, error) {
	var configs Configs

	loadFile(&configs)

	configs.DB = configs.DBConfig.pgConfig()

	return configs, nil
}

func loadFile(configs *Configs) {
	f, err := os.Open(configPath)
	if err != nil {
		handleError(err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(configs)
	handleError(err)
}

func (cfg *DBConfig) pgConfig() *sqlx.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=public sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	dbConn := sqlx.MustConnect("pgx", connStr)

	return dbConn
}

func handleError(err error) {
	if err == nil {
		return
	}
	log.Println("config loading error", err)
	panic(err)
}
