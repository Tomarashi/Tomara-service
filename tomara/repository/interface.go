package repository

type IRepository interface {
}

type ITomaraRepository interface {
	IRepository
	CheckDatabase() (bool, error)
	DBType() string
	GetWordsStartsWith(substring string, limit int) (error, []string)
	GetWordsContains(substring string, limit int) (error, []string)
}
