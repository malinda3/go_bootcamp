package mymath

import (
	"math"
	"sort"
)

func Calculate(stat int, s []int) interface{} {
	if len(s) == 0 {
		return "Error: input array is empty"
	}
	switch stat {
	case 0:
		return Mean(s)
	case 1:
		return Median(s)
	case 2:
		return Mode(s)
	case 3:
		return SD(s)
	default:
		return "Error"
	}
}

func Mean(s []int) float64 {
	res := 0.0
	for i := 0; i < len(s); i++ {
		res += float64(s[i])
	}
	return res / float64(len(s))
}

func Median(s []int) float64 {
    sort.Ints(s)
    if len(s)%2 != 0 {
        return float64(s[len(s)/2])
    }
    return float64(s[len(s)/2-1]+s[len(s)/2]) / 2.0
}

func Mode(s []int) int {
    frequency := make(map[int]int)
    maxCount := 0
    res := s[0]

    for i := 0; i < len(s); i++ {
		num := s[i]
		frequency[num]++
		if frequency[num] > maxCount {
			maxCount = frequency[num]
			res = num
		}
	}
	// for _, num := range s {
    //     frequency[num]++
    //     if frequency[num] > maxCount {
    //         maxCount = frequency[num]
    //         res = num
    //     }
    // }
    return res
}



func SD(s []int) float64 {
	if len(s) <= 1 {
		return 0.0
	}
    meanVal := Mean(s)
    var sumSquares float64

    for _, num := range s {
        diff := float64(num) - meanVal
        sumSquares += diff * diff
    }

    sd := math.Sqrt(sumSquares / float64(len(s) - 1))
    return sd
}

