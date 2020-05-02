package FirstBadVersion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	start := 0
	target := n / 2
	end := n

	for {
		// badならこれより手前に切り替わりがある
		if isBadVersion(target) {
			// 終了条件は、target = end
			if target == end {
				return target - 1
			}
			end = target
			target = start + (end-start)/2
		} else {
			// 終了条件は、target = start
			if target == start {
				return target + 1
			}
			start = target
			target = start + (end-start)/2
		}
	}
}

// mock
func isBadVersion(version int) bool {
	return true
}
