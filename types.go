package quranquiz

type Verse struct {
	ID              int    `json:"id"`
	VerseNumber     int    `json:"verse_number"`
	VerseKey        string `json:"verse_key"`
	JuzNumber       int    `json:"juz_number"`
	HizbNumber      int    `json:"hizb_number"`
	RubElHizbNumber int    `json:"rub_el_hizb_number"`
	RukuNumber      int    `json:"ruku_number"`
	ManzilNumber    int    `json:"manzil_number"`
	SajdahNumber    int    `json:"sajdah_number"`
	ChapterID       int    `json:"chapter_id"`
	TextImlaei      string `json:"text_imlaei"`
	TextUthmani     string `json:"text_uthmani"`
	PageNumber      int    `json:"page_number"`
}

type Surah struct {
	Name    string `json:"name"`
	Chapter int    `json:"chapter"`
	Verse   int    `json:"verse"`
}

type Quran struct {
	Verses []Verse `json:"verses"`
	Surah  []Surah `json:"surah"`
}

type Option struct {
	Text      string `json:"text"`
	Correct   bool   `json:"correct"`
	ChapterID int    `json:"chapter_id,omitempty"`
	VerseID   int    `json:"verse_id,omitempty"`
}

type Quiz struct {
	Question      string `json:"question"`
	ChapterID     int    `json:"chapter_id,omitempty"`
	VerseID       int    `json:"verse_id,omitempty"`
	CorrectAnswer int    `json:"correct_answer"`

	Options []Option `json:"options"`
}
