package main

import "fmt"
// https://pintia.cn/problem-sets/1211841066264109056/problems/1211848231062290433
func main() {
	var l int
	fmt.Scan(&l)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		var c int
		fmt.Scan(&c)
		arr[i] = c
	}
	thissum, maxsum := 0, 0
	start, end := 0, 0
	// start 应该要和 end 绑定 这种flag的判断方式不适合
	//for i := 0; i < l; i++ {
	//	thissum += arr[i]
	//	if !minExist && thissum >= 0  {
	//		start = i
	//		minExist = true
	//	}
	//	if  thissum > maxsum {
	//		maxsum = thissum
	//		end = i
	//	}else if thissum < 0 {
	//		thissum = 0
	//		minExist = false
	//	}
	//}
	// 这里有两种方式 一个用来计数的count--> end - count +1 = start
	/* 另外一种方式
	if thissum > maxsum {
			maxsum = thissum
			end = i
			start = leftSwitch
		}
	if(ThisSum<0){
		ThisSum = 0;
		leftSwitch = i+1;
	}
	*/
	count, flag := 0, 0
	for i := 0; i < l; i ++ {
		thissum += arr[i]
		count ++
		if thissum > maxsum {
			maxsum = thissum
			end = i
			start = end - count + 1
		}else if thissum < 0 {
			thissum = 0
			count = 0
		}else if thissum == 0 {
			if maxsum == 0 && flag == 0 {
				flag = 1
				start = i
				end = i
			}
		}
	}
	if maxsum == 0 && flag == 0 {
		fmt.Println(maxsum, arr[0], arr[l-1])
	}else{
		fmt.Println(maxsum, arr[start], arr[end])
	}

}

/*
7
0  1  2  3  5  0  -7
10
-10 1 2 3 4 -5 -23 3 7 -21
3
-2 -3 -1  // 0 -2 -1
3
0 -3 3
 */
