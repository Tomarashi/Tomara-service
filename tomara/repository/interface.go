package repository

type IRepository interface {
}

type ITomaraRepository interface {
	IRepository
	GetWordsStartsWith(substring string, limit int) (error, []string)
	GetWordsContains(substring string, limit int) (error, []string)
}
