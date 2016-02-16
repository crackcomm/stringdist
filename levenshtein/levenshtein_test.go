// The MIT License (MIT)
//
// Copyright (c) 2014 Vlad-Doru Ion
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package levenshtein

import (
	"testing"
)

func TestDistanceThreshold(t *testing.T) {
	threshold := 2
	var testCases = []struct {
		source   string
		target   string
		distance int
		within   bool
	}{
		{"", "aa", 2, true},
		{"a", "aa", 1, true},
		{"a", "aaa", 2, true},
		{"", "", 0, true},
		{"a", "bcaa", -1, false},
		{"aaa", "aba", 1, true},
		{"aaa", "abcd", -1, false},
		{"a", "a", 0, true},
		{"ab", "aabc", 2, true},
		{"abc", "", -1, false},
		{"aa", "a", 1, true},
		{"aaaaa", "a", -1, false},
		{"informatica", "fmi unibuc", -1, false},
	}
	for _, testCase := range testCases {
		distance, within := DistanceThreshold(testCase.source, testCase.target, threshold)

		if within != testCase.within {
			t.Log("Distance between",
				testCase.source,
				"and",
				testCase.target,
				"computed as",
				within,
				", should be",
				testCase.within)
			t.Error("Failed to compute threshold properly")
		}

		if within && distance != testCase.distance {
			t.Log("Distance between",
				testCase.source,
				"and",
				testCase.target,
				"computed as",
				distance,
				", should be",
				testCase.distance)
			t.Error("Failed to compute proper Levenshtein Distance")
		}
	}
}
