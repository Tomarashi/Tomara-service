package repository

import (
	"bufio"
	"index/suffixarray"
	"os"
	"sort"
	"strings"
)

type FakeRepository struct {
	data *suffixarray.Index
}

func (f FakeRepository) extractWords(indexes []int) []string {
	result := make([]string, 0, len(indexes))
	bytesData := f.data.Bytes()
	for _, index := range indexes {
		left := index
		for bytesData[left] != 0 {
			left -= 1
		}
		right := left + 1
		for bytesData[right] != 0 {
			right += 1
		}
		result = append(result, string(bytesData[left+1:right]))
	}
	return result
}

func (f FakeRepository) GetWordsStartsWith(substring string, limit int) []string {
	substringBytes := append([]byte{0}, []byte(substring)...)
	preResult := f.extractWords(f.data.Lookup(substringBytes, -1))
	sort.Strings(preResult)
	if limit < 0 || limit > len(preResult) {
		limit = len(preResult)
	}
	return preResult[:limit]
}

func (f FakeRepository) GetWordsContains(substring string, limit int) []string {
	preResult := f.extractWords(f.data.Lookup([]byte(substring), -1))
	postResult := make([]string, 0, len(preResult))
	set := make(map[string]bool)
	for _, res := range preResult {
		_, ok := set[res]
		if !ok {
			set[res] = true
			postResult = append(postResult, res)
		}
	}
	sort.Strings(postResult)
	if limit < 0 || limit > len(preResult) {
		limit = len(preResult)
	}
	return preResult[:limit]
}

func MakeFakeRepository(words []string) *FakeRepository {
	data := make([]byte, 0, 1024)
	for _, word := range words {
		data = append(data, 0)
		data = append(data, []byte(word)...)
	}
	data = append(data, 0)
	return &FakeRepository{data: suffixarray.New(data)}
}

func MakeFakeRepositoryFromFile(filePath string) *FakeRepository {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		result = append(result, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
	return MakeFakeRepository(result)
}
