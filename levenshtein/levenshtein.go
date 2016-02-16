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

import "math"

// DistanceThreshold computes the Levenshtein distance between two strings
// if and only if it is smaller than a specific threshold
//
// Arugments:
// source, target (string): thw two strings to compute the distance for
// threshold (int): the threshold of the Levenshtein distance
//
// Returns: (int, bool) Levenshtein distance and if it is lower than the
// threshold. The first value is valid iff the second one is true.
func DistanceThreshold(source, target string, threshold int) (int, bool) {
	sourceLen := len(source)
	targetLen := len(target)

	if threshold == 0 && source != target {
		return 0, false
	}

	if dist := int(math.Abs(float64(sourceLen - targetLen))); dist > threshold {
		return 0, false
	}

	if sourceLen > targetLen {
		source, target = target, source
		sourceLen, targetLen = targetLen, sourceLen
	}

	diff := targetLen - sourceLen

	v0, v1 := make([]int, targetLen+1), make([]int, targetLen+1)

	for i := 0; i <= targetLen; i++ {
		v0[i] = i
	}

	cost, lower := 0, 0 // Lower bound at each step
	for i := 0; i < sourceLen; i++ {
		start, stop := max(0, i-threshold), min(targetLen, i+diff+threshold)
		if start == 0 {
			v1[start] = i + 1
		} else {
			cost = 0
			if source[i-1] != target[start-1] {
				cost = 1
			}
			v1[start] = min(v0[start]+1, v0[start-1]+cost)
		}
		lower = v1[start]
		for j := start; j < stop-1; j++ {
			cost = 0
			if source[i] != target[j] {
				cost = 1
			}
			v1[j+1] = min3(v1[j]+1, v0[j+1]+1, v0[j]+cost)
			lower = min(v1[j+1], lower)
		}
		cost = 0
		if source[i] != target[stop-1] {
			cost = 1
		}
		v1[stop] = min(v1[stop-1]+1, v0[stop-1]+cost)
		lower = min(v1[stop], lower)
		// If the lower bound is higher than the threshold we return false
		if lower > threshold {
			return -1, false
		}
		v0, v1 = v1, v0
	}

	return v0[targetLen], v0[targetLen] <= threshold
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func min3(x, y, z int) int {
	if x < y && x < z {
		return x
	}
	if y < z {
		return y
	}
	return z
}
