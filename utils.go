package quranquiz

import (
	"errors"
	"math/rand"
)

func randIntSlice(haystack []int) int {
	return haystack[rand.Intn(len(haystack))]
}

func randVerselice(haystack []Verse) Verse {
	return haystack[rand.Intn(len(haystack))]
}

func pickOne() Verse {
	return QURAN.Verses[rand.Intn(len(QURAN.Verses))]
}

func pickMany(amount int) (v []Verse) {
	existsVerseID := make(map[int]bool)

	safeLoop := 0

	for i := 0; i < amount; i++ {
		verse := pickOne()
		if _, ok := existsVerseID[verse.ID]; !ok {
			v = append(v, pickOne())
		}
		safeLoop++

		if safeLoop > amount && safeLoop > 1000 {
			return v
		}
	}

	return v
}

func randomChapter() int {
	min := 1
	max := 114
	return rand.Intn(max-min+1) + min
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func inSliceInt(haystack []int, needle int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

func getManyVerseIdBySurahId(chapters []int) (v []Verse) {

	for i := range QURAN.Verses {
		verse := QURAN.Verses[i]

		if inSliceInt(chapters, verse.ChapterID) {
			v = append(v, verse)
		}

	}

	return v
}

func getSurahByVerseID(verseID int) (int, string, error) {
	var verse Verse

	for _, v := range QURAN.Verses {
		if v.ID == verseID {
			verse = v
			break
		}
	}

	if verse.ID == 0 {
		return 0, "", errors.New("ayat tidak ditemukan")
	}

	return verse.ChapterID, GetSurahName(verse.ChapterID), nil
}

func CheckSurahName(chapterID int) (string, error) {
	if chapterID < 1 || chapterID > 114 {
		return "", errors.New("nomor surah tidak valid")
	}

	return QURAN.Surah[chapterID-1].Name, nil
}

func GetSurahName(chapterID int) string {
	surahName, _ := CheckSurahName(chapterID)
	return surahName
}

func uniqueInts(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func randomOptions(opt []Option) {
	rand.Shuffle(len(opt), func(i, j int) { opt[i], opt[j] = opt[j], opt[i] })
}
