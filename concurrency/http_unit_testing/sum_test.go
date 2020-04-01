package sum

import "testing"

func TestInts(t *testing.T) {
	
	tt := []struct{
		name string 
		numbers []int
		sum int
	}{
		{"one to five" , []int{1,2,3,4,5},15},
		{"no numbers", nil , 0},
		{"aon and minis one" , []int{1,-1},0},
	}
	
	for _ , tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := ints(tc.numbers)
			if s != tc.sum {
				t.Fatalf("sum of %v should be %v; got %v",tc.name , tc.sum , s)
			}
		})
	}
	
}
