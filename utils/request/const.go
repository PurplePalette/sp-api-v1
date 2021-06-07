package request

const SPLIT_QUERY = ":"
const SPLIT_DIFFICULTY = "-"

type SearchSortType string

const (
	SORT_UPDATED_DATE SearchSortType = "updatedDate"
	SORT_CREATED_DATE SearchSortType = "createdDate"
	SORT_NAME         SearchSortType = "name"
	SORT_DIFFICULTY   SearchSortType = "rating"
	SORT_NOTES        SearchSortType = "notes"
)

var supportedSorts map[string]SearchSortType = map[string]SearchSortType{
	"date":         SORT_UPDATED_DATE,
	"updateddate":  SORT_UPDATED_DATE,
	"updated_date": SORT_UPDATED_DATE,
	"createddate":  SORT_CREATED_DATE,
	"created_date": SORT_CREATED_DATE,
	"name":         SORT_NAME,
	"t":            SORT_NAME,
	"title":        SORT_NAME,
	"d":            SORT_DIFFICULTY,
	"difficulty":   SORT_DIFFICULTY,
	"l":            SORT_DIFFICULTY,
	"level":        SORT_DIFFICULTY,
	"r":            SORT_DIFFICULTY,
	"rating":       SORT_DIFFICULTY,
	"n":            SORT_NOTES,
	"notes":        SORT_NOTES,
}

type SearchSortOrder int

const (
	ORDER_DESC SearchSortOrder = 0
	ORDER_ASC  SearchSortOrder = 1
)

var supportedOrders map[string]SearchSortOrder = map[string]SearchSortOrder{
	"desc": ORDER_DESC,
	"d":    ORDER_DESC,
	"asc":  ORDER_ASC,
	"a":    ORDER_ASC,
}

type SearchFilterGenre string

const (
	GENRE_ALL      SearchFilterGenre = ""
	GENRE_GENERAL  SearchFilterGenre = "general"
	GENRE_JPOP     SearchFilterGenre = "jpop"
	GENRE_ANIME    SearchFilterGenre = "anime"
	GENRE_VOCALOID SearchFilterGenre = "vocaloid"
)

var supportedGenres map[string]SearchFilterGenre = map[string]SearchFilterGenre{
	"all":      GENRE_ALL,
	"general":  GENRE_GENERAL,
	"g":        GENRE_GENERAL,
	"jpop":     GENRE_JPOP,
	"j":        GENRE_JPOP,
	"anime":    GENRE_ANIME,
	"a":        GENRE_ANIME,
	"vocaloid": GENRE_VOCALOID,
	"v":        GENRE_VOCALOID,
}

type SearchFilterDifficulty [2]int

type SearchFilter struct {
	Difficulty  SearchFilterDifficulty
	Genre       SearchFilterGenre
	UserId      string
	Keyword     string
	Public      bool
	ForcePublic bool
}

type SearchQuery struct {
	Sort   SearchSortType
	Order  SearchSortOrder
	Filter SearchFilter
}
