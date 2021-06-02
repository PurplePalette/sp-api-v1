package potato

import (
	"encoding/json"
	"errors"
	"math"
	"sort"
	"strings"

	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
)

type Cache struct {
	Data map[string]interface{}
}

func (c *Cache) IsExist(key string) bool {
	_, ok := c.Data[key]
	return ok
}

func (c *Cache) Set(key string, value interface{}) {
	c.Data[key] = value
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.Data[key]
	if !ok {
		return nil, errors.New("could not find the data for specified key")
	}
	return value, nil
}

func (c *Cache) Add(key string, data interface{}) error {
	_, ok := c.Data[key]
	if ok {
		return errors.New("specified key already exists")
	}
	c.Data[key] = data
	return nil
}

func (c *Cache) Remove(key string, data interface{}) error {
	_, ok := c.Data[key]
	if ok {
		return errors.New("specified key already exists")
	}
	delete(c.Data, key)
	return nil
}

func (c *Cache) IsOwnerMatch(key string, userId string) (bool, error) {
	v, ok := c.Data[key]
	if !ok {
		return false, errors.New("specified key was not exists")
	}
	switch d := v.(type) {
	case Background:
		return d.UserId == userId, nil
	case Effect:
		return d.UserId == userId, nil
	case Engine:
		return d.UserId == userId, nil
	case Particle:
		return d.UserId == userId, nil
	case Skin:
		return d.UserId == userId, nil
	case Level:
		return d.UserId == userId, nil
	default:
		return false, nil
	}
}

func (c *Cache) Pages() int32 {
	dataLen := float64(len(c.Data))
	return int32(math.Ceil(dataLen / 20))
}

func Paginate(pageNum int32, pageSize int, sliceLength int) (int, int) {
	start := int(pageNum) * pageSize

	if start > sliceLength {
		start = sliceLength
	}

	end := start + pageSize
	if end > sliceLength {
		end = sliceLength
	}

	return start, end
}

func (c *Cache) GetPage(pageId int32, searchQuery request.SearchQuery) ([]byte, error) {
	// Convert map to slice
	var dataList DataList
	for _, value := range c.Data {
		dataList = append(dataList, value)
	}
	// Sort slice using query
	orderQuery := searchQuery.Order
	switch searchQuery.Sort {
	case request.SORT_CREATED_DATE:
		if orderQuery == request.ORDER_ASC {
			sort.Sort(sort.Reverse(ByCreatedTime{dataList}))
		} else {
			sort.Sort(ByCreatedTime{dataList})
		}
	case request.SORT_UPDATED_DATE:
		if orderQuery == request.ORDER_ASC {
			sort.Sort(sort.Reverse(ByUpdatedTime{dataList}))
		} else {
			sort.Sort(ByUpdatedTime{dataList})
		}
	case request.SORT_DIFFICULTY:
		if orderQuery == request.ORDER_ASC {
			sort.Sort(sort.Reverse(ByDifficulty{dataList}))
		} else {
			sort.Sort(ByDifficulty{dataList})
		}
	case request.SORT_NOTES:
		if orderQuery == request.ORDER_ASC {
			sort.Sort(sort.Reverse(ByNotes{dataList}))
		} else {
			sort.Sort(ByNotes{dataList})
		}
	}
	var filteredList DataList
	for _, value := range dataList {
		switch v := value.(type) {
		case Level:
			minD := int32(searchQuery.Filter.Difficulty[0])
			maxD := int32(searchQuery.Filter.Difficulty[1])
			ratingMatched := v.Rating >= minD && v.Rating <= maxD
			genreMatched := strings.Contains(v.Genre, string(searchQuery.Filter.Genre))
			userMatched := strings.Contains(v.UserId, searchQuery.Filter.UserId)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if ratingMatched && genreMatched && userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Background:
			userMatched := strings.Contains(v.UserId, searchQuery.Filter.UserId)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Effect:
			userMatched := strings.Contains(v.UserId, searchQuery.Filter.UserId)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Engine:
			userMatched := strings.Contains(v.UserId, searchQuery.Filter.UserId)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Particle:
			userMatched := strings.Contains(v.UserId, searchQuery.Filter.UserId)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Skin:
			userMatched := strings.Contains(v.UserId, searchQuery.Filter.UserId)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		}
	}
	start, end := Paginate(pageId, 20, len(filteredList))
	jsonData, err := json.Marshal(filteredList[start:end])
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
