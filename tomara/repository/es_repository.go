package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"golang.org/x/net/context"
	"io"
)

const (
	anyCharRegex     = ".*"
	wordGeoFieldName = "word_geo"
)

type ElasticSearchRepository struct {
	client *elasticsearch.Client
}

type buildQueryOptions struct {
	fieldName         string
	subValue          string
	useOnlyStartsWith bool
	resultSize        int
}

func buildQuery(opts buildQueryOptions) (error, []byte) {
	var buf bytes.Buffer
	var valueRegex string
	if opts.useOnlyStartsWith {
		valueRegex = opts.subValue + anyCharRegex
	} else {
		valueRegex = anyCharRegex + opts.subValue + anyCharRegex
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"regexp": map[string]interface{}{
				opts.fieldName: map[string]interface{}{
					"value": valueRegex,
					"flags": "ALL",
				},
			},
		},
		"sort": []map[string]interface{}{
			{
				"frequency": map[string]string{
					"order": "desc",
				},
			},
		},
		"size": opts.resultSize,
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return err, nil
	}
	return nil, buf.Bytes()
}

func (e *ElasticSearchRepository) CheckDatabase() (bool, error) {
	_, err := e.client.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *ElasticSearchRepository) DBType() string {
	return "ElasticSearchRepository"
}

func (e *ElasticSearchRepository) makeRequest(queryBytes []byte) (error, []string) {
	castErrorFmt := func(fieldName string) error {
		return fmt.Errorf("can't cast '%s'", fieldName)
	}

	request := esapi.SearchRequest{
		Index:      []string{"words"},
		Body:       bytes.NewReader(queryBytes),
		FilterPath: []string{"hits.hits._source.word_geo", "took"},
	}
	response, err := request.Do(context.Background(), e.client)
	if err != nil {
		return err, nil
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err, nil
	}
	err = response.Body.Close()
	if err != nil {
		return err, nil
	}

	bodyMap := make(map[string]interface{})
	if err := json.Unmarshal(bodyBytes, &bodyMap); err != nil {
		return err, nil
	}

	hitsMap, castOk := bodyMap["hits"].(map[string]interface{})
	if !castOk {
		return castErrorFmt("hits"), nil
	}

	sourceList, castOk := hitsMap["hits"].([]interface{})
	if !castOk {
		return castErrorFmt("hits.hits"), nil
	}

	result := make([]string, 0, len(sourceList))
	for i, sourceElem := range sourceList {
		sourceMap, castOk := sourceElem.(map[string]interface{})
		if !castOk {
			return castErrorFmt(fmt.Sprintf("hits.hits[%d]", i)), nil
		}
		val, castOk := sourceMap["_source"].(map[string]interface{})
		if !castOk {
			return castErrorFmt(fmt.Sprintf("hits.hits[%d]_source", i)), nil
		}
		result = append(result, val["word_geo"].(string))
	}
	return nil, result
}

func (e *ElasticSearchRepository) findWords(substring string, limit int, useOnlyStartsWith bool) (error, []string) {
	err, queryBytes := buildQuery(buildQueryOptions{
		fieldName:         wordGeoFieldName,
		resultSize:        limit,
		subValue:          substring,
		useOnlyStartsWith: useOnlyStartsWith,
	})
	if err != nil {
		return err, nil
	}

	err, result := e.makeRequest(queryBytes)
	if err != nil {
		return err, nil
	}
	return nil, result
}

func (e *ElasticSearchRepository) GetWordsStartsWith(substring string, limit int) (error, []string) {
	return e.findWords(substring, limit, true)
}

func (e *ElasticSearchRepository) GetWordsContains(substring string, limit int) (error, []string) {
	return e.findWords(substring, limit, false)
}

func MakeElasticSearchRepository() *ElasticSearchRepository {
	repository := &ElasticSearchRepository{}
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	repository.client = client
	return repository
}
