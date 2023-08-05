package repository

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearchRepository struct {
	Client elasticsearch.Client
}

func createQuery(fieldName string, resultSize int) string {
	return fmt.Sprintf(`{
		"query": {
			"regexp": {
				"%s": {
					"flags": "ALL",
					"value": "ჰ.*"
				}
			}
		},
		"size": %d,
		"sort": [{
			"frequency": {
				"order": "desc"
			}
		}]
	}`, fieldName, resultSize)
}

func (e ElasticSearchRepository) GetWordsStartsWith(substring string, limit int) []string {
	//TODO implement me
	panic(NotImplementedErrorString)
}

func (e ElasticSearchRepository) GetWordsContains(substring string, limit int) []string {
	//TODO implement me
	panic(NotImplementedErrorString)
}
