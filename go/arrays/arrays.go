package main

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {
    var sum_all []int
    for _,slice := range slices {
        sum := 0
        for _,num := range slice {
            sum += num
        }
        sum_all = append(sum_all, sum)
    }
    return sum_all
}
