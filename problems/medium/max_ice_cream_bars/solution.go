package maxicecreambars

/*
https://leetcode.com/problems/maximum-ice-cream-bars/?envType=daily-question&envId=2026-06-21
*/

func brought(cost int, coins int, freq int) int {
	if cost <= coins && freq > 0 {
		if freq*cost > coins {
			freq = coins / cost
		}
		return freq
	}
	return 0
}

func maxIceCream(costs []int, coins int) int {
	max_cost := 0

	for _, cost := range costs {
		if cost > max_cost {
			max_cost = cost
		}
	}

	freq_array := make([]int, max_cost+1)
	for _, cost := range costs {
		freq_array[cost]++
	}

	count := 0
	for cost := 1; cost <= max_cost; cost++ {
		if freq_array[cost] == 0 {
			continue
		}

		if cost <= coins && freq_array[cost] > 0 {
			brought_count := brought(cost, coins, freq_array[cost])
			count += brought_count
			coins -= cost * brought_count
		}

		if cost > coins {
			break
		}
	}

	return count
}
