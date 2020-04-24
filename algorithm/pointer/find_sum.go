package pointer

import "math"

/*
* 剑指offer 41
* 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。
* 如果有多对数字的和等于s，输出任意一对即可。
 */

func FindSumInArray(array []int, sum int) []int {

	head := 0
	tail := len(array) - 1

	if len(array) < 2 {
		return nil
	}

	//创建两个指针

	for array[head]+array[tail] != sum {
		// 如果两个数相加小于sum，说明需要将前指针向后移
		if array[head]+array[tail] < sum {
			head++
			if head == tail {
				return nil
			}
		}

		if array[head]+array[tail] > sum {
			tail--
			if head == tail {
				return nil
			}
		}
	}

	return []int{array[head], array[tail]}

}

/*
 * 输入一个正数s，一个连续数组，打印出所有和为s可能的数组
 * 例如输入15，[1,2,3,4,5,6,7,8,9,10]，所以结果打印出3个连续序列1～5、4～6和7～8。
 */

func FindSequenceArray(array []int, sum int) (output [][]int) {
	// 是否是空数组
	if array == nil {
		return nil
	}

	// 如果只有一个元素
	if len(array) == 1 {
		if sum == array[0] {
			output = append(output, []int{array[0]})
			return
		}

		return nil
	}

	head := 0
	tail := 1

	getSum := func(array []int) (output int) {
		for _, v := range array {
			output += v
		}
		return
	}

	for {

		if head == tail {
			if getSum(array[head:tail]) == sum {
				output = append(output, array[head:tail])
			}
			return output
		}

		if getSum(array[head:tail]) < sum {
			if tail == len(array)-1 {
				return output
			}

			tail++
		}

		if getSum(array[head:tail]) > sum {
			if head == tail-1 {
				return output
			}

			head++
		}

		if getSum(array[head:tail]) == sum {
			output = append(output, array[head:tail])

			if tail == len(array)-1 {
				return
			}

			tail++
		}

	}

}

/*
 * 输入一个正数s，打印出所有和为s的连续正数序列（至少含有两个数）。
 * 例如输入15，由于1＋2＋3＋4＋5＝4＋5＋6＝7＋8＝15，所以结果打印出3个连续序列1～5、4～6和7～8。
 */

func FindSequenceArrayFromNumber(num int) (output [][]int) {
	if num < 3 {
		return nil
	}

	head := 1
	tail := 2

	getSum := func(begin, end int) (output int) {

		for i := begin; i <= end; i++ {
			output += i
		}

		return output

	}

	generateArray := func(begin, end int) (array []int) {
		for i := begin; i <= end; i++ {
			array = append(array, i)
		}

		return array
	}

	for math.Ceil(float64(1+num)/2) > float64(tail) {

		if getSum(head, tail) < num {
			tail++
		}

		if getSum(head, tail) > num {
			head++
		}

		if getSum(head, tail) == num {
			output = append(output, generateArray(head, tail))
			tail++
		}

	}

	return

}
