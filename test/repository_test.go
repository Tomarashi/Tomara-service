package test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"tomara-service/tomara/repository"
)

func TestFakeRepository(t *testing.T) {
	var result []string
	var testWord string

	fakeRepository := repository.MakeFakeRepositoryFromFile(TestDataFilePath)

	result = fakeRepository.GetWordsContains("ა", 200)
	assert.Equal(t, 200, len(result))

	result = fakeRepository.GetWordsContains("აა", math.MaxInt)
	assert.Equal(t, 26, len(result))

	testWord = "დადებითი"
	result = fakeRepository.GetWordsContains(testWord, math.MaxInt)
	assert.Equal(t, []string{testWord}, result)

	result = fakeRepository.GetWordsStartsWith("აა", math.MaxInt)
	assert.Equal(t, 0, len(result))

	result = fakeRepository.GetWordsStartsWith("ბი", math.MaxInt)
	assert.Equal(t, 24, len(result))

	result = fakeRepository.GetWordsStartsWith("ერთ", 1)
	assert.Equal(t, []string{"ერთ"}, result)

	result = fakeRepository.GetWordsStartsWith("ერთ", math.MaxInt)
	assert.True(t, contains(result, "ერთ-ერთი"))
}
