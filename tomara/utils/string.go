package utils

const (
	GeoLetters      = "აბგდევზთიკლმნოპჟრსტუფქღყშჩცძწჭხჯჰ"
	geoAsEngLetters = "abgdevzTiklmnopJrstufqRySCcZwWxjh"
)

var GeoToEngMap map[rune]rune

func init() {
	GeoToEngMap = make(map[rune]rune)

	geoLettersArray := []rune(GeoLetters)
	geoAsEngLettersArray := []rune(geoAsEngLetters)
	for i, c := range geoLettersArray {
		GeoToEngMap[c] = geoAsEngLettersArray[i]
	}
}

func GeoWordToEng(geoWord string) string {
	result := make([]rune, 0, len(geoWord))
	for _, c := range geoWord {
		result = append(result, GeoToEngMap[c])
	}
	return string(result)
}
