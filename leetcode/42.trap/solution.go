package _2_trap

func trap(height []int) int {
	upStack := make([]int, 0)
	result := 0
	for i := 0; i < len(height); i++ {
		if i == 0 {
			goto hok
		}
		if height[i] > height[i-1] {
			if i == 1 {
				goto hok
			}
			for len(upStack) != 0 {
				pre := upStack[len(upStack)-1]
				size := (i - pre - 1) * calculate(height, pre, i)
				result += size
				if height[pre] > height[i] {
					break
				}
				upStack = upStack[:len(upStack)-1]
			}
		}
	hok:
		if i == len(height)-1 {
			break
		}
		if height[i] > height[i+1] {
			upStack = append(upStack, i)
		}
	}
	return result
}

func calculate(height []int, s, e int) int {
	max := -1
	for i := s + 1; i < e; i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	min := -1
	if height[s] > height[e] {
		min = height[e]
	} else {
		min = height[s]
	}
	return min - max
}
