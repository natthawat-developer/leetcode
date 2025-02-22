package problems


func reverse(x int) int {
    sign := 1
    if x < 0 {
        sign = -1
        x = -x
    }

    reversed := 0
    for x > 0 {
        reversed = reversed*10 + x%10
        x /= 10
    }

    reversed *= sign

    if reversed < -2147483648 || reversed > 2147483647 {
        return 0
    }

    return reversed
}

