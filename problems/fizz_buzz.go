package problems

import "fmt"

func FizzBuzz(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0{
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
            result = append(result, "Fizz")
        }else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, fmt.Sprintf("%d", i))
		}
	}

	return result
}

// import "strconv"

// func FizzBuzz(n int) []string {
//     result := make([]string, n)

//     for i := 1; i <= n; i++ {
//         if i%3 == 0 && i%5 == 0 {
//             result[i-1] = "FizzBuzz"
//         } else if i%3 == 0 {
//             result[i-1] = "Fizz"
//         } else if i%5 == 0 {
//             result[i-1] = "Buzz"
//         } else {
//             result[i-1] = strconv.Itoa(i) // แปลงตัวเลขเป็น string
//         }
//     }

//     return result
// }