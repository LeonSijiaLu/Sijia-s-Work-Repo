package main

import "fmt"

func calculate_no_mask_day(days int, masks_left int) int {
	masks_left = masks_left + days
	enough_masks_left := false
	masks := 0
	got_mask_days := 0
	for !enough_masks_left {
		masks = masks + 5
		got_mask_days = got_mask_days + 1
		if masks == masks_left || masks > masks_left {
			enough_masks_left = true
		}
	}
	return got_mask_days
}

func main() {
	fmt.Println(calculate_no_mask_day(10, 20))
	fmt.Println(calculate_no_mask_day(20, 60))
}
