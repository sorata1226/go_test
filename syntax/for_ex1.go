package main

import "fmt"

func main() {
	// 0 から 99 までの数字を順に追加したスライスを作成する

	nums := make([]int, 0, 100)

	for i := 0; i<100; i++{

		//fmt.Printf("nums = %d, e = %d\n",nums,e)
		nums = append(nums, i)
	}
		fmt.Printf("nums = %v", nums)

	//fmt.Printf("nums = %v\n", nums)

	// 最後の要素から逆順にスライスの内容を出力する

	for j := len(nums) -1; j >= 0; j--{
		fmt.Printf("%v",nums[j])
	}

}
