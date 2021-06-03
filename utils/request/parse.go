package request

import (
	"strconv"
	"strings"
)

func parseDifficulty(difficulty string) [2]int {
	difficulties := strings.Split(difficulty, SPLIT_DIFFICULTY)
	if len(difficulties) == 1 {
		if num, err := strconv.Atoi(difficulty); err == nil {
			return [2]int{num, num}
		}
	} else if difficulties[0] == "" || difficulties[1] == "" {
		if strings.HasPrefix(difficulty, SPLIT_DIFFICULTY) {
			if num, err := strconv.Atoi(difficulties[1]); err == nil {
				return [2]int{0, num}
			}
		}
		if strings.HasSuffix(difficulty, SPLIT_DIFFICULTY) {
			if num, err := strconv.Atoi(difficulties[0]); err == nil {
				return [2]int{num, 100}
			}
		}
	} else {
		if numA, err := strconv.Atoi(difficulties[0]); err == nil {
			if numB, err2 := strconv.Atoi(difficulties[1]); err2 == nil {
				if numA < numB {
					return [2]int{numA, numB}
				} else {
					return [2]int{numB, numA}
				}
			}
		}
	}
	return [2]int{0, 100}
}

// ParseSearchQuery parses the sonolus keywords to query
func ParseSearchQuery(keywords string) SearchQuery {
	resp := SearchQuery{
		Filter: SearchFilter{
			Difficulty: [2]int{0, 100},
			Public:     true,
		},
	}
	var originalKeywords []string
	queries := strings.Split(keywords, " ")
	for _, word := range queries {
		query := strings.Split(word, SPLIT_QUERY)
		if len(query) != 2 {
			originalKeywords = append(originalKeywords, word)
			continue
		}
		queryKey := strings.ToLower(query[0])
		queryValue := strings.ToLower(query[1])
		switch queryKey {
		case "s", "sort":
			if v, ok := supportedSorts[queryValue]; ok {
				resp.Sort = v
			}
		case "o", "order":
			if v, ok := supportedOrders[queryValue]; ok {
				resp.Order = v
			}
		case "g", "genre":
			if v, ok := supportedGenres[queryValue]; ok {
				resp.Filter.Genre = v
			}
		case "d", "difficulty":
			resp.Filter.Difficulty = parseDifficulty(queryValue)
		case "u", "user":
			resp.Filter.UserId = queryValue
		}
	}
	resp.Filter.Keyword = strings.Join(originalKeywords, " ")
	return resp
}
