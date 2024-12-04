package mymath_test

import (
	"main/mymath"
	"testing"
)

func TestCalc0(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{3.0, 3.0, 1, 1.58}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcNoInput(t *testing.T) {
	numbers := []int{}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expectedResults := []any{"Error: input array is empty", "Error: input array is empty", "Error: input array is empty", "Error: input array is empty"}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case string:
			if v != expectedResults[i].(string) {
				t.Errorf("Expected %s: %v, got: %s", statName[i], expectedResults[i], v)
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcSingleElement(t *testing.T) {
	numbers := []int{10}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{10.0, 10.0, 10, 0.0}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcAllSameValues(t *testing.T) {
	numbers := []int{7, 7, 7, 7}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{7.0, 7.0, 7, 0.0}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcTwoElements(t *testing.T) {
	numbers := []int{1, 2}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{1.5, 1.5, 1, 0.71}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}
func TestCalcNegativeOddArray(t *testing.T) {
	numbers := []int{-1, -3, -5, -7, -9}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{-5.0, -5.0, -9, 3.16}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcAllNegativeValues(t *testing.T) {
	numbers := []int{-10, -20, -30, -40, -50}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{-30.0, -30.0, -50, 15.81}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcSmallDecimalValues(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{4.0, 4.0, 1, 2.16}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcLargerArrayWithRepeatingValues(t *testing.T) {
	numbers := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{3.0, 3.0, 1, 1.46}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcUnsortedArray(t *testing.T) {
	numbers := []int{5, 1, 4, 2, 3}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{3.0, 3.0, 1, 1.58}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcEmptyStringInput(t *testing.T) {
	numbers := []int{}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{"Error: input array is empty", "Error: input array is empty", "Error: input array is empty", "Error: input array is empty"}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case string:
			if v != expected[i].(string) {
				t.Errorf("Expected %s: %v, got: %s", statName[i], expected[i], v)
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}
func TestCalcNegativeValues(t *testing.T) {
	numbers := []int{-1, -2, -3, -4, -5}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{-3.0, -3.0, -5, 1.58}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcLargeArray(t *testing.T) {
	numbers := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = i % 10 
	}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{4.5, 4.5, 0, 2.87}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcZeros(t *testing.T) {
	numbers := []int{0, 0, 0, 1, 2, 3}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{1.0, 0.5, 0, 1.26}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcMixedValues(t *testing.T) {
	numbers := []int{-3, 5, 0, -2, 8, -3, 5, 5}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{1.88, 2.5, 5, 4.35}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcLargePositiveValues(t *testing.T) {
	numbers := []int{1000000, 2000000, 3000000, 4000000, 5000000}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{3000000.0, 3000000.0, 1000000, 1581138.83}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcNilValues(t *testing.T) {
	numbers := []int{0, 0, 0, 0}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{0.0, 0.0, 0, 0.0}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcSingleNegativeElement(t *testing.T) {
	numbers := []int{-10}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{-10.0, -10.0, -10, 0.0}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

func TestCalcArrayWithDuplicates(t *testing.T) {
	numbers := []int{1, 1, 2, 2, 3, 3, 3}
	stats := []int{0, 1, 2, 3}
	statName := []string{"Mean", "Median", "Mode", "SD"}
	expected := []any{2.14, 2.0, 3, 0.90}

	for i, stat := range stats {
		result := mymath.Calculate(stat, numbers)

		switch v := result.(type) {
		case int:
			if expected[i] != v {
				t.Errorf("Expected %s: %v, got: %d", statName[i], expected[i], v)
			}
		case float64:
			if expectedVal, ok := expected[i].(float64); ok {
				if diff := v - expectedVal; diff < -0.01 || diff > 0.01 {
					t.Errorf("Expected %s: %.2f, got: %.2f", statName[i], expectedVal, v)
				}
			} else {
				t.Errorf("Unexpected type for %s: got float64, expected int", statName[i])
			}
		default:
			t.Errorf("Unexpected result type for %s: %T", statName[i], result)
		}
	}
}

//test only math func

func TestMeanWithOddNumberOfElements(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	stat := 0
	expected := 3.0

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected Mean: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestMeanWithEvenNumberOfElements(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	stat := 0
	expected := 2.5

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected Mean: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestMedianWithOddNumberOfElements(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	stat := 1
	expected := 3.0

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected Median: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestMedianWithEvenNumberOfElements(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	stat := 1
	expected := 2.5

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected Median: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestModeWithMultipleFrequentValues(t *testing.T) {
	numbers := []int{1, 2, 2, 3, 3, 4}
	stat := 2
	expected := 2

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(int); ok {
		if v != expected {
			t.Errorf("Expected Mode: %d, got: %d", expected, v)
		}
	} else {
		t.Errorf("Expected result type: int, got: %T", result)
	}
}

func TestStandardDeviationWithEmptyInput(t *testing.T) {
	numbers := []int{}
	stat := 3
	expected := "Error: input array is empty"

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(string); ok {
		if v != expected {
			t.Errorf("Expected Error: %v, got: %s", expected, v)
		}
	} else {
		t.Errorf("Expected result type: string, got: %T", result)
	}
}

func TestStandardDeviationWithSingleValue(t *testing.T) {
	numbers := []int{1}
	stat := 3
	expected := 0.0

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected SD: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestModeWithSingleElement(t *testing.T) {
	numbers := []int{5}
	stat := 2
	expected := 5

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(int); ok {
		if v != expected {
			t.Errorf("Expected Mode: %d, got: %d", expected, v)
		}
	} else {
		t.Errorf("Expected result type: int, got: %T", result)
	}
}

func TestSDWithNegativeNumbers(t *testing.T) {
	numbers := []int{-1, -2, -3, -4, -5}
	stat := 3
	expected := 1.58

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if diff := v - expected; diff < -0.01 || diff > 0.01 {
			t.Errorf("Expected SD: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestMeanWithLargeValues(t *testing.T) {
	numbers := []int{1000000, 2000000, 3000000, 4000000, 5000000}
	stat := 0
	expected := 3000000.0

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected Mean: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestMedianWithLargeArray(t *testing.T) {
	numbers := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		numbers[i] = i
	}
	stat := 1
	expected := 499.5

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected Median: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestModeWithSingleOccurrence(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	stat := 2
	expected := 1

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(int); ok {
		if v != expected {
			t.Errorf("Expected Mode: %d, got: %d", expected, v)
		}
	} else {
		t.Errorf("Expected result type: int, got: %T", result)
	}
}

func TestSDWithLargeNumbers(t *testing.T) {
	numbers := []int{1000000, 2000000, 3000000, 4000000, 5000000}
	stat := 3
	expected := 1581138.83

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if diff := v - expected; diff < -0.01 || diff > 0.01 {
			t.Errorf("Expected SD: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestMeanWithNegativeNumbers(t *testing.T) {
	numbers := []int{-1, -2, -3, -4, -5}
	stat := 0
	expected := -3.0

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(float64); ok {
		if v != expected {
			t.Errorf("Expected Mean: %.2f, got: %.2f", expected, v)
		}
	} else {
		t.Errorf("Expected result type: float64, got: %T", result)
	}
}

func TestMedianWithEmptyArray(t *testing.T) {
	numbers := []int{}
	stat := 1
	expected := "Error: input array is empty"

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(string); ok {
		if v != expected {
			t.Errorf("Expected Error: %v, got: %s", expected, v)
		}
	} else {
		t.Errorf("Expected result type: string, got: %T", result)
	}
}

func TestModeWithNoRepeatingValues(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	stat := 2
	expected := 1

	result := mymath.Calculate(stat, numbers)
	if v, ok := result.(int); ok {
		if v != expected {
			t.Errorf("Expected Mode: %d, got: %d", expected, v)
		}
	} else {
		t.Errorf("Expected result type: int, got: %T", result)
	}
}




