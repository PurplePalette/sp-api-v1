package request_test

import (
	"testing"

	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
	"github.com/stretchr/testify/assert"
)

func TestParseEmptyText(t *testing.T) {
	want := request.SearchQuery{}
	actual := request.ParseSearchQuery("")
	assert.Equal(t, want, actual)
}

func TestParseKeywordOne(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Keyword: "NoPoi!",
		},
	}
	actual := request.ParseSearchQuery("NoPoi!")
	assert.Equal(t, want, actual)
}

func TestParseKeywordTwo(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Keyword: "Daydream cafe",
		},
	}
	actual := request.ParseSearchQuery("Daydream cafe")
	assert.Equal(t, want, actual)
}

func TestParseSortName(t *testing.T) {
	want := request.SearchQuery{
		Sort: request.SORT_NAME,
	}
	actual := request.ParseSearchQuery("sort:name")
	assert.Equal(t, want, actual)
}

func TestParseSortNameNoValue(t *testing.T) {
	want := request.SearchQuery{
		Sort: request.SORT_UPDATED_DATE,
	}
	actual := request.ParseSearchQuery("sort:")
	assert.Equal(t, want, actual)
}

func TestParseSortNameShortHand(t *testing.T) {
	want := request.SearchQuery{
		Sort: request.SORT_NAME,
	}
	actual := request.ParseSearchQuery("s:t")
	assert.Equal(t, want, actual)
}

func TestParseFilterGenre(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Genre: request.GENRE_ANIME,
		},
	}
	actual := request.ParseSearchQuery("genre:anime")
	assert.Equal(t, want, actual)
}

func TestParseFilterGenreShortHand(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Genre: request.GENRE_ANIME,
		},
	}
	actual := request.ParseSearchQuery("g:a")
	assert.Equal(t, want, actual)
}

func TestParseKeywordOneAndSortName(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Keyword: "NoPoi!",
		},
		Sort: request.SORT_NAME,
	}
	actual := request.ParseSearchQuery("NoPoi! sort:name")
	assert.Equal(t, want, actual)
}

func TestParseKeywordTwoAndSortName(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Keyword: "Daydream cafe",
		},
		Sort: request.SORT_NAME,
	}
	actual := request.ParseSearchQuery("Daydream cafe sort:name")
	assert.Equal(t, want, actual)
}

func TestParseDifficultyNoValue(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Difficulty: [2]int{0, 100},
		},
	}
	actual := request.ParseSearchQuery("difficulty:-")
	assert.Equal(t, want, actual)
}

func TestParseDifficultyOneValue(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Difficulty: [2]int{10, 10},
		},
	}
	actual := request.ParseSearchQuery("difficulty:10")
	assert.Equal(t, want, actual)
}

func TestParseDifficultyOneValueUpperThan(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Difficulty: [2]int{10, 100},
		},
	}
	actual := request.ParseSearchQuery("difficulty:10-")
	assert.Equal(t, want, actual)
}

func TestParseDifficultyOneValueLowerThan(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Difficulty: [2]int{0, 50},
		},
	}
	actual := request.ParseSearchQuery("difficulty:-50")
	assert.Equal(t, want, actual)
}

func TestParseDifficultyTwoValue(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Difficulty: [2]int{10, 50},
		},
	}
	actual := request.ParseSearchQuery("difficulty:10-50")
	assert.Equal(t, want, actual)
}

func TestParseDifficultyTwoValueFlipped(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Difficulty: [2]int{10, 50},
		},
	}
	actual := request.ParseSearchQuery("difficulty:50-10")
	assert.Equal(t, want, actual)
}

func TestParseMixed(t *testing.T) {
	want := request.SearchQuery{
		Sort: request.SORT_NOTES,
		Filter: request.SearchFilter{
			Keyword:    "Daydream cafe ～チノver.～",
			Difficulty: [2]int{10, 30},
			Genre:      request.GENRE_ANIME,
			UserId:     "domao",
		},
	}
	actual := request.ParseSearchQuery("Daydream cafe difficulty:10-30 genre:anime sort:notes user:domao ～チノver.～")
	assert.Equal(t, want, actual)
}

func TestParseMixedShortHand(t *testing.T) {
	want := request.SearchQuery{
		Sort: request.SORT_NOTES,
		Filter: request.SearchFilter{
			Keyword:    "Daydream cafe ～チノver.～",
			Difficulty: [2]int{10, 30},
			Genre:      request.GENRE_ANIME,
			UserId:     "domao",
		},
	}
	actual := request.ParseSearchQuery("Daydream cafe d:10-30 g:a s:n u:domao ～チノver.～")
	assert.Equal(t, want, actual)
}

func TestParseDifficultyTwice(t *testing.T) {
	want := request.SearchQuery{
		Filter: request.SearchFilter{
			Difficulty: [2]int{5, 10},
		},
	}
	actual := request.ParseSearchQuery("difficulty:10-30 difficulty:5-10")
	assert.Equal(t, want, actual)
}
