package anna

/* import "test/base" */

func replaceSignArray(n int, a []string) [51]string {
	var newSignArr [51]string

	for i := 0; i < len(a); i++ {
		sign := a[i]
		if i < n {
			j := i
			j += len(a)
			j -= n
			newSignArr[j] = sign
		} else {
			j := i
			j -= n
			newSignArr[j] = sign
		}
	}

	return newSignArr
}
