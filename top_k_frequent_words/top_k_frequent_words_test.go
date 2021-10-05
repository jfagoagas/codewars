package topKFrequent

import (
	"testing"
)

var validateTopKFrequent = []struct {
	name     string
	words    []string
	key      int
	expected []string
}{
	{"Test 1", []string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
	{"Test 2", []string{"i", "love", "leetcode", "i", "love", "coding"}, 3, []string{"i", "love", "coding"}},
	{"Test 3", []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	{"Test 4", []string{"i", "love", "leetcode", "i", "love", "coding"}, 3, []string{"i", "love", "coding"}},
	{"Test 5", []string{"plpaboutit", "jnoqzdute", "sfvkdqf", "mjc", "nkpllqzjzp", "foqqenbey", "ssnanizsav", "nkpllqzjzp", "sfvkdqf", "isnjmy", "pnqsz", "hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "cyxwlef", "hhqpvvt", "fvvdtpnzx", "plpaboutit", "sfvkdqf", "mjc", "fvvdtpnzx", "bwumsj", "foqqenbey", "isnjmy", "nkpllqzjzp", "hhqpvvt", "foqqenbey", "fvvdtpnzx", "bwumsj", "hhqpvvt", "fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "foqqenbey", "jnoqzdute", "foqqenbey", "hhqpvvt", "ssnanizsav", "mjc", "foqqenbey", "bwumsj", "ssnanizsav", "fvvdtpnzx", "nkpllqzjzp", "jkqonvenhx", "hhqpvvt", "mjc", "isnjmy", "bwumsj", "pnqsz", "hhqpvvt", "nkpllqzjzp", "jnoqzdute", "pnqsz", "nkpllqzjzp", "jnoqzdute", "foqqenbey", "nkpllqzjzp", "hhqpvvt", "fvvdtpnzx", "plpaboutit", "jnoqzdute", "sfvkdqf", "fvvdtpnzx", "jkqonvenhx", "jnoqzdute", "nkpllqzjzp", "jnoqzdute", "fvvdtpnzx", "jkqonvenhx", "hhqpvvt", "isnjmy", "jkqonvenhx", "ssnanizsav", "jnoqzdute", "jkqonvenhx", "fvvdtpnzx", "hhqpvvt", "bwumsj", "nkpllqzjzp", "bwumsj", "jkqonvenhx", "jnoqzdute", "pnqsz", "foqqenbey", "sfvkdqf", "sfvkdqf"}, 1, []string{"fvvdtpnzx"}},
}

func TestTopKFrequent(t *testing.T) {
	for _, tt := range validateTopKFrequent {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.words, tt.key)
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Expected: %v, got: %v", tt.expected, got)
				}
			}
		})
	}
}
