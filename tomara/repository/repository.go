package repository

import (
	"fmt"
)

func logRepoError(dbType string, err error) {
	fmt.Printf("Can't connect to database '%s'. Error: %s\n", dbType, err.Error())
}

func GetFirstValidRepository() ITomaraRepository {
	esRepo := MakeElasticSearchRepository()
	if valid, err := esRepo.CheckDatabase(); valid {
		return esRepo
	} else {
		logRepoError("ElasticSearch", err)
	}
	mySqlRepo := MakeMySqlRepositoryDefaultConfig()
	if valid, err := mySqlRepo.CheckDatabase(); valid {
		return mySqlRepo
	} else {
		logRepoError("MySql", err)
	}
	return MakeFakeRepositoryFromFile("data/dev-words.txt")
}
