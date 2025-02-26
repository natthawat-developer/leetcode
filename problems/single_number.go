package problems

func SingleNumber(nums []int) int {

	hash := make(map[int]int)

	for _, num := range nums {
		hash[num]++
	}

	for key, value := range hash {
		if value == 1 {
			return key
		}
	}

	return 0
}


// func SingleNumber(nums []int) int {
//     result := 0
//     for _, num := range nums {
//         result ^= num // ใช้ XOR เพื่อลบค่าที่ซ้ำออก
//     }
//     return result
// }