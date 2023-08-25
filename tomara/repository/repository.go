package repository

func GetFirstValidRepository() ITomaraRepository {
	esRepo := MakeElasticSearchRepository()
	if valid, _ := esRepo.CheckDatabase(); valid {
		return esRepo
	}
	mySqlRepo := MakeMySqlRepositoryDefaultConfig()
	if valid, _ := mySqlRepo.CheckDatabase(); valid {
		return mySqlRepo
	}
	return MakeFakeRepositoryFromFile("data/dev-words.txt")
}
