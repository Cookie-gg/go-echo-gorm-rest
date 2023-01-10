package util

import "encoding/json"

func GetJsonKyes(j []byte) []string {
	// a map container to decode the JSON structure into
	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s := range c {
		k[i] = s
		i++
	}

	return k
}

func Filter[T any](array []T, f func(T) bool) (result []T) {
	for _, value := range array {
		if f(value) {
			result = append(result, value)
		}
	}
	return
}

func Map[T, T2 any](array []T, f func(T) T2) (result []T2) {
	for _, value := range array {
		result = append(result, f(value))
	}
	return
}
