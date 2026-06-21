package jumpgamevii

// https://leetcode.com/problems/jump-game-vii/description/?envType=daily-question&envId=2026-05-25

/*
1. i + minJump <= j <= min(i + maxJump, s.length - 1)
2. s[j] == '0'

"011010"
"012345"
minJump = 2, maxJump = 3

i = 0

c10 = 0 + 2 <= 1 <= min(0 + 3, 6 - 1) -> false

c20:  0 + 2 <= 2 <= min(0 + 3, 6 - 1) -> true
c21: s[2] == '1' -> false

c30: 0 + 2 <= 3 <= min(0 + 3, 6 - 1) -> true
c31: s[3] == '0' -> true

i = 3

brute force way to do it


*/

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)

	if s[n-1] != '0' {
		return false
	}

	queue := make(chan int, n)
	queue <- 0
	max_visited := 0
    for len(queue) > 0 {
		current := <-queue

		if current == n - 1  {
			return true
		}

		start_range := max(current + minJump, max_visited + 1)
		end_range := min(current + maxJump, n - 1)

		for j := start_range; j <= end_range; j++ {
			if s[j] == '0' {
				queue <- j
			}
		}

		max_visited = max(max_visited, end_range)
	}
	return false
}
