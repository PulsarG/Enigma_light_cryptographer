package anna

/* import "test/base" */

func replaceSignArray(n int, a []string) []string {
	var newSignArr []string

	for i := 0; i < len(a); i++ {
		sign := a[i]
		if i < n {
			j := i
			j += len(a)
			j -= n
			newSignArr = append(newSignArr, sign)
			/* newSignArr[j] = sign */
		} else {
			j := i
			j -= n
			newSignArr = append(newSignArr, sign)
			/* newSignArr[j] = sign */
		}
	}

	return newSignArr
}

func replaceRotor(n int, a []int) []int {
	var newSignArr []int

	for i := 0; i < len(a); i++ {
		sign := a[i]
		if i < n {
			j := i
			j += len(a)
			j -= n
			newSignArr = append(newSignArr, sign)
			/* newSignArr[j] = sign */
		} else {
			j := i
			j -= n
			newSignArr = append(newSignArr, sign)
			/* newSignArr[j] = sign */
		}
	}

	return newSignArr
}
