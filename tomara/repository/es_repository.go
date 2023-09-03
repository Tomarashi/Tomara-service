package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"golang.org/x/net/context"
	"io"
	"strings"
)

const (
	anyCharRegex     = ".*"
	wordGeoFieldName = "word_geo"
)

var (
	requestIndex      []string
	requestFilterPath []string
)

func init() {
	requestIndex = []string{"words"}
	requestFilterPath = []string{
		"hits.hits._source.word_geo",
		"hits.total.value",
		"took",
	}
}

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
	extractValue := func(values map[string]interface{}, fields string) (error, interface{}) {
		var castOk bool
		fieldsArr := strings.Split(fields, ".")
		currVal := values
		for i, field := range fieldsArr {
			if i == len(fieldsArr)-1 {
				break
			}
			currVal, castOk = currVal[field].(map[string]interface{})
			if !castOk {
				fieldName := strings.Join(fieldsArr[0:i+1], ".")
				return fmt.Errorf("can't cast '%s'", fieldName), nil
			}
		}
		return nil, currVal[fieldsArr[len(fieldsArr)-1]]
	}

	request := esapi.SearchRequest{
		Index:      requestIndex,
		Body:       bytes.NewReader(queryBytes),
		FilterPath: requestFilterPath,
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

	var sourceList []interface{}
	if err, val := extractValue(bodyMap, "hits.hits"); err == nil {
		if val == nil {
			return nil, []string{}
		} else {
			sourceList = val.([]interface{})
		}
	} else {
		return err, nil
	}

	result := make([]string, 0, len(sourceList))
	for _, sourceElem := range sourceList {
		err, val := extractValue(sourceElem.(map[string]interface{}), "_source.word_geo")
		if err != nil {
			return err, nil
		}
		result = append(result, val.(string))
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
