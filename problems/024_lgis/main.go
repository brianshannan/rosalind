package main

import (
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/brianshannan/rosalind/utils"
)

func main() {
	flag.Parse()
	data, err := utils.ReadFileFromArgStripTrailingNewline()
	if err != nil {
		panic(err)
	}

	sequenceLine := strings.Split(string(data), "\n")[1]
	sequenceStrings := strings.Split(sequenceLine, " ")
	sequence := make([]int, len(sequenceStrings))
	for i, str := range sequenceStrings {
		sequence[i], err = strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
	}

	lis := longestIncreasingSubsequence(sequence)
	fmt.Printf("%v\n", strings.Join(utils.SliceMap(lis, strconv.Itoa), " "))

	// longest decreasing is longest increasing with negated inputs
	negativeSequence := negateSlice(sequence)
	lds := longestIncreasingSubsequence(negativeSequence)
	lds = negateSlice(lds)
	fmt.Printf("%v\n", strings.Join(utils.SliceMap(lds, strconv.Itoa), " "))
}

func negateSlice(vals []int) []int {
	negativeVals := make([]int, len(vals))
	for i, value := range vals {
		negativeVals[i] = -1 * value
	}
	return negativeVals
}

// https://en.wikipedia.org/wiki/Longest_increasing_subsequence
func longestIncreasingSubsequence(sequence []int) []int {
	// indexes into sequence for our current candidate subsequences.
	// each i for bestSubsequenceIdxs[i] tracks the "best" possible subsequence
	// of length i.
	var bestSubsequenceIdxs []int

	// tracks the previous index into sequence for the longest
	// subsequence this index is part of. This allows us to
	// walk backward to find the subsequence chain
	parents := make([]int, len(sequence))
	for i := range parents {
		parents[i] = -1
	}

	for i := 0; i < len(sequence); i++ {
		// if we're the first element, we have to be the biggest thing
		// we don't have a parent as we're the first element of the subsequence
		if len(bestSubsequenceIdxs) == 0 {
			bestSubsequenceIdxs = append(bestSubsequenceIdxs, i)
			continue
		}

		// if we're bigger than the current last element, we just extend the subsequence
		// and make the previous element in the subsequence our parent
		currentSubsequenceEndIdx := bestSubsequenceIdxs[len(bestSubsequenceIdxs)-1]
		if sequence[i] > sequence[currentSubsequenceEndIdx] {
			parents[i] = currentSubsequenceEndIdx
			bestSubsequenceIdxs = append(bestSubsequenceIdxs, i)
			continue
		}

		// we can't just extend the subsequence, go back
		// and check if we can make a new "branch" of the subsequence
		// as we made a better subsequence of length idxToReplace than we previously had
		idxToReplace := findSmallestLargerValue(sequence, bestSubsequenceIdxs, i)
		bestSubsequenceIdxs[idxToReplace] = i

		// if we replace the smallest element, we started a new subsequence
		// and don't have a parent
		// otherwise the previous element in the subsequence is our parent
		if idxToReplace > 0 {
			parents[i] = bestSubsequenceIdxs[idxToReplace-1]
		}
	}

	var lis []int
	for currParent := bestSubsequenceIdxs[len(bestSubsequenceIdxs)-1]; currParent > -1; currParent = parents[currParent] {
		lis = append(lis, sequence[currParent])
	}

	slices.Reverse(lis)
	return lis
}

// findSmallestLargerValue finds the smallest value in subsequenceIdxs that valueIdx is smaller than or equal to
// Let x = sequence[subsequenceIdx[i]] for some i. We want to find the smallest i
// such that sequence[valueIdx] <= x.
// subsequenceIdxs should be sorted such that sequence[subsequenceIdxs[i]] < sequence[subsequenceIdxs[j]] if i < j
func findSmallestLargerValue(sequence []int, subsequenceIdxs []int, valueIdx int) int {
	low := 0
	high := len(subsequenceIdxs) - 1
	for high > low {
		mid := low + (high-low)/2
		midValue := sequence[subsequenceIdxs[mid]]
		targetValue := sequence[valueIdx]

		if midValue == targetValue {
			// We found a value equal to ourself, so return that index
			return mid
		} else if midValue < targetValue {
			// we're bigger than our mid point,
			// so the element we're smaller than is to our right
			low = mid + 1
		} else {
			// we're smaller than our mid point,
			// so the element we're smaller than is to our left or the mid point
			high = mid
		}
	}
	return high
}
