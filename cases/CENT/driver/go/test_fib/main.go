package main

func GenFibonacci(limit int) []int {
	//result := 0
	//var array []int
	var array []int
	for i := 0; i < limit; i++ {
		array = append(array, fibonacci(i))
		//result = fibonacci(i)
		//array = append(array, result)
		//fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	return array
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
