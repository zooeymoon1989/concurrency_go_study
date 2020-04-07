package algorithm

import "testing"

func TestFindAllDuplicatesInAnArray(t *testing.T)  {

	result := 11
	outputResult := 0
	output := FindAllDuplicatesInAnArray(10,2,5,10,9,1,1,4,3,7)

	for _ , v := range output {
		outputResult += v
	}

	if result != outputResult {
		t.Errorf("sum of array should be %d , got %d" ,result, outputResult)
	}
}
