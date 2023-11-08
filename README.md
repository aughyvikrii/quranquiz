# Quran Quiz

repository ini adalah versi bahasa pemrograman go dari repository https://github.com/zakiego/quran-quiz oleh [@zakiego](https://github.com/zakiego)

## Installation

```bash
go get github.com/aughyvikrii/quranquiz
```

## Example

```bash
import (
  "fmt"

  "github.com/aughyvikrii/quranquiz"
)

func main() {

  /** generate 1 random quiz guess by surah from any chapter */
  quizzes, err := quranquiz.GuessSurah(1)

  /** generate 2 random quiz guess by surah chapter (1, 2, 3, 4) */
  quizzes, err := quranquiz.GuessSurah(2, 1, 2, 3, 4)
}
```


## Note
Mohon maaf jika ada kesalahan pada library ini, silahkan buat PR untuk memperbaiki library