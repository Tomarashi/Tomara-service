package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"time"
	"tomara-service/tomara/configs"
)

const (
	responseLimit = 500
)

var config MySqlRepositoryConfig

type MySqlRepository struct {
	database *sql.DB
}

type MySqlRepositoryConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func init() {
	err := yaml.Unmarshal(configs.MySqlConfigYamlString, &config)
	if err != nil {
		panic(err)
	}
}

func (m MySqlRepository) formatQueryString(
	substring string, limit int, onlyStartsWith bool,
) string {
	var searchTypeString string
	if onlyStartsWith {
		searchTypeString = ""
	} else {
		searchTypeString = "%"
	}
	return fmt.Sprintf(
		`SELECT word_geo FROM words WHERE word_geo LIKE "%s%s%%" ORDER BY frequency DESC LIMIT %d`,
		searchTypeString, substring, limit,
	)
}

func (m MySqlRepository) getWordsAny(substring string, limit int, onlyStartsWith bool) []string {
	if limit < 0 || limit > responseLimit {
		limit = responseLimit
	}
	queryString := m.formatQueryString(substring, limit, onlyStartsWith)
	queryResult, err := m.database.Query(queryString)
	if err != nil {
		panic(err)
	}
	result := make([]string, 0, limit)
	for queryResult.Next() {
		var row string
		if err := queryResult.Scan(&row); err != nil {
			panic(err)
		}
		result = append(result, row)
	}
	if err := queryResult.Close(); err != nil {
		panic(err)
	}
	return result
}

func (m MySqlRepository) GetWordsStartsWith(substring string, limit int) []string {
	return m.getWordsAny(substring, limit, true)
}

func (m MySqlRepository) GetWordsContains(substring string, limit int) []string {
	return m.getWordsAny(substring, limit, false)
}

func MakeMySqlRepository(username string, password string, databaseName string) *MySqlRepository {
	sourceName := fmt.Sprintf("%s:%s@/%s", username, password, databaseName)
	database, err := sql.Open("mysql", sourceName)
	if err != nil {
		panic(err)
	}
	database.SetConnMaxLifetime(time.Minute)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)
	return &MySqlRepository{database: database}
}

func MakeMySqlRepositoryDefaultConfig() *MySqlRepository {
	return MakeMySqlRepository(config.User, config.Password, config.Database)
}
