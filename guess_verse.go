package quranquiz

import (
	"errors"
)

/**
* amount = amount of quiz to generate
* chapters = chapter to generate or get verse from (if empty, lib will generate random chapter)
 */

func GuessVerse(amount int, chapters ...int) (quizzes []Quiz, err error) {
	if len(chapters) == 0 {
		for i := 0; i < 4; i++ {
			chapters = append(chapters, randomChapter())
		}
	} else if len(chapters) < 4 {
		return nil, errors.New("minimal pilih 4 surah yang berbeda")
	} else {
		chaptersUnique := uniqueInts(chapters)
		if len(chaptersUnique) < 4 {
			return nil, errors.New("minimal pilih 4 surah yang berbeda")
		}
	}

	min, max := minMax(chapters)

	if min < 1 || max > 114 {
		return nil, errors.New("pilih surah antara 1 s/d 114")
	}

	verses := getManyVerseIdBySurahId(chapters)

	createQuiz := func() (quiz Quiz, err error) {

		var (
			questionVerse Verse
			surah         Surah
		)

		for {
			questionVerse = randVerselice(verses)
			surah = QURAN.Surah[questionVerse.ChapterID-1]

			if questionVerse.VerseNumber != surah.Verse { // if the last verse, then find another
				break
			}
		}

		nextVerse := QURAN.Verses[questionVerse.ID]

		options := []Option{
			{
				Text:    GetSurahName(nextVerse.ChapterID),
				Correct: true,
				VerseID: nextVerse.ID,
			},
		}

		for i := len(options); i < 4; i++ {
			option := randVerselice(verses)

			if option.ID != nextVerse.ID && option.ID != questionVerse.ID {
				options = append(options, Option{
					Text:    option.TextUthmani,
					Correct: false,
					VerseID: option.ID,
				})
			}
		}

		randomOptions(options)

		quiz = Quiz{
			Question: questionVerse.TextUthmani,
			VerseID:  questionVerse.ID,
			Options:  options,
		}

		return quiz, nil
	}

	for i := 0; i < amount; i++ {
		quiz, err := createQuiz()
		if err != nil {
			return nil, err
		}

		quizzes = append(quizzes, quiz)
	}

	return quizzes, nil
}
