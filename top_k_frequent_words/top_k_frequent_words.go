package topKFrequent

import (
	"sort"
)

/*
	Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
	Output: ["the","is","sunny","day"]
	Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
	with the number of occurrence being 4, 3, 2 and 1 respectively.
*/
type wordFrecuency struct {
	word      string
	frequency int
}

func topKFrequent(words []string, k int) []string {
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	var topK []wordFrecuency
	for word, frequency := range wordCount {
		topK = append(topK, wordFrecuency{word, frequency})
	}

	// sort slice frequency in ascending order by their frequency
	sort.Slice(topK, func(i, j int) bool {
		var compare bool
		if topK[i].frequency == topK[j].frequency {
			compare = topK[i].word < topK[j].word
		} else {
			compare = topK[i].frequency > topK[j].frequency
		}
		return compare
	})
	// fmt.Println(topK)

	var output []string
	for i := 0; i < k; i++ {
		output = append(output, topK[i].word)
	}

	// fmt.Println(output)
	return output
}
