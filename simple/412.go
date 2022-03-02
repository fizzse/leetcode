package main

import "strconv"

func fizzBuzz(n int) []string {
	var ret []string
	m := map[int]string{
		1: "Fizz",
		2: "Buzz",
		3: "FizzBuzz",
	}
	nums := []int{3, 5}
	for i := 1; i <= n; i++ {
		cnt := 0
		for index, v := range nums {
			if i%v == 0 {
				cnt += index + 1
			}
		}

		if m[cnt] != "" {
			ret = append(ret, m[cnt])
			continue
		}
		ret = append(ret, strconv.Itoa(i))
		// 0 1,2 3
	}

	return ret
}
