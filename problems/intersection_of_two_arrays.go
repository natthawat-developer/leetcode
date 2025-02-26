package problems

//import "fmt"

func Intersect(nums1 []int, nums2 []int) []int {

	result := make([]int,0)
	for _,num1 := range nums1 {
		
		for i,num2 := range nums2{
	
			if num1 == num2{
				nums2 = append(nums2[:i], nums2[i+1:]...)
				result = append(result, num1)
				break
			}
		
		}
	}
	return result
}

// func Intersect(nums1 []int, nums2 []int) []int {
//     freq := make(map[int]int)
//     result := []int{}

//     // นับจำนวนครั้งที่ตัวเลขปรากฏใน nums1
//     for _, num := range nums1 {
//         freq[num]++
//     }

//     // ตรวจสอบ nums2 และลดค่าจำนวนครั้งที่พบ
//     for _, num := range nums2 {
//         if freq[num] > 0 {
//             result = append(result, num)
//             freq[num]--
//         }
//     }

//     return result
// }