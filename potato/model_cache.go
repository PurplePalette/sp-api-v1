package potato

import (
	"encoding/json"
	"errors"
	"math"
	"sort"
	"strings"

	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
)

// Cache is data store struct for storeing data
type Cache struct {
	Data map[string]interface{}
}

// IsExist check the specified key is exists in cache
func (c *Cache) IsExist(key string) bool {
	_, ok := c.Data[key]
	return ok
}

// Set sets the value to cache with using specified key
func (c *Cache) Set(key string, value interface{}) {
	c.Data[key] = value
}

// Get gets value from cache with using specified key.
// It returns error if key does not exist in cache.
func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.Data[key]
	if !ok {
		return nil, errors.New("could not find the data for specified key")
	}
	return value, nil
}

// Add adds the value to cache with using specified key.
// It returns error if key already exists.
func (c *Cache) Add(key string, value interface{}) error {
	_, ok := c.Data[key]
	if ok {
		return errors.New("specified key already exists")
	}
	c.Data[key] = value
	return nil
}

// Remove removes data from cache that has specified key.
func (c *Cache) Remove(key string, data interface{}) error {
	_, ok := c.Data[key]
	if !ok {
		return errors.New("specified key not exists")
	}
	delete(c.Data, key)
	return nil
}

// IsOwnerMatch check the owner of specified content is same as specified userID.
// It returns error if specified key was not existed.
func (c *Cache) IsOwnerMatch(key string, userID string) (bool, error) {
	v, ok := c.Data[key]
	if !ok {
		return false, errors.New("specified key was not exists")
	}
	switch d := v.(type) {
	case Background:
		return d.UserID == userID, nil
	case Effect:
		return d.UserID == userID, nil
	case Engine:
		return d.UserID == userID, nil
	case Particle:
		return d.UserID == userID, nil
	case Skin:
		return d.UserID == userID, nil
	case Level:
		return d.UserID == userID, nil
	default:
		return false, nil
	}
}

// Pages returns the length of cache divided by 20 and also roundup proceeded
func (c *Cache) Pages() int32 {
	dataLen := float64(len(c.Data))
	return int32(math.Ceil(dataLen / 20))
}

// paginate return start of page, and end of page, used for getting specified page
func paginate(pageNum int32, pageSize int, sliceLength int) (int, int) {
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

// GetPage gets specified page from cache datalist.
// 20 contents will be returned as json bytes array if succeed.
// It returns error if it couldn't marshal to json bytes array.
func (c *Cache) GetPage(pageID int32, searchQuery request.SearchQuery) ([]byte, error) {
	// Convert map to slice
	var dataList DataList
	for _, value := range c.Data {
		dataList = append(dataList, value)
	}
	// Sort slice using query
	orderQuery := searchQuery.Order
	switch searchQuery.Sort {
	case request.SORT_CREATED_DATE:
		if orderQuery == request.ORDER_DESC {
			sort.Sort(sort.Reverse(ByCreatedTime{dataList}))
		} else {
			sort.Sort(ByCreatedTime{dataList})
		}
	case request.SORT_UPDATED_DATE:
		if orderQuery == request.ORDER_DESC {
			sort.Sort(sort.Reverse(ByUpdatedTime{dataList}))
		} else {
			sort.Sort(ByUpdatedTime{dataList})
		}
	case request.SORT_DIFFICULTY:
		if orderQuery == request.ORDER_DESC {
			sort.Sort(sort.Reverse(ByDifficulty{dataList}))
		} else {
			sort.Sort(ByDifficulty{dataList})
		}
	case request.SORT_NOTES:
		if orderQuery == request.ORDER_DESC {
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
			userMatched := strings.Contains(v.UserID, searchQuery.Filter.UserID)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			publicMatched := v.Public == searchQuery.Filter.Public || searchQuery.Filter.ForcePublic
			if ratingMatched && genreMatched && userMatched && keywordMatched && publicMatched {
				filteredList = append(filteredList, value)
			}
		case Background:
			userMatched := strings.Contains(v.UserID, searchQuery.Filter.UserID)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Effect:
			userMatched := strings.Contains(v.UserID, searchQuery.Filter.UserID)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Engine:
			userMatched := strings.Contains(v.UserID, searchQuery.Filter.UserID)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Particle:
			userMatched := strings.Contains(v.UserID, searchQuery.Filter.UserID)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case Skin:
			userMatched := strings.Contains(v.UserID, searchQuery.Filter.UserID)
			keywordMatched := strings.Contains(v.Title, searchQuery.Filter.Keyword)
			if userMatched && keywordMatched {
				filteredList = append(filteredList, value)
			}
		case User:
			return nil, errors.New("user data pagination is not supported yet")
		}
	}
	start, end := paginate(pageID, 20, len(filteredList))
	jsonData, err := json.Marshal(filteredList[start:end])
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
