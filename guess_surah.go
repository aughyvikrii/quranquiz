package quranquiz

import (
	"errors"
)

/**
* amount = amount of quiz to generate
* chapters = chapter to generate or get verse from (if empty, lib will generate random chapter)
 */

func GuessSurah(amount int, chapters ...int) (quizzes []Quiz, err error) {
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
		question := randVerselice(verses)

		options := []Option{
			{
				Text:      GetSurahName(question.ChapterID),
				Correct:   true,
				ChapterID: question.ChapterID,
			},
		}

		for i := range chapters {
			chapterID := chapters[i]
			if chapterID == question.ChapterID {
				continue
			}

			options = append(options, Option{
				Text:      GetSurahName(chapterID),
				Correct:   false,
				ChapterID: chapterID,
			})
		}

		randomOptions(options)

		correctAnswer := 0

		for i := range options {
			if options[i].Correct {
				correctAnswer = i
				break
			}
		}
		quiz = Quiz{
			Question:      question.TextUthmani,
			VerseID:       question.ID,
			ChapterID:     question.ChapterID,
			CorrectAnswer: correctAnswer,
			Options:       options,
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
