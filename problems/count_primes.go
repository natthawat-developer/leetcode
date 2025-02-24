package problems

import "math"

func CountPrimes(n int) int {
    if n < 2 {
        return 0
    }

    count := 0
    for i := 2; i < n; i++ {
        isPrime := true
        for j := 2; j <= int(math.Sqrt(float64(i))); j++ { // เช็คจาก 2 ถึง sqrt(i)
            if i%j == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            count++
        }
    }

    return count
}
