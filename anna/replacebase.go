package anna

import "test/base"

func replaceSignArray(n int) [51]string {
	var newSignArr [51]string

	for i := 0; i < len(base.SignsArray); i++ {
		sign := base.SignsArray[i]
		if i < n {
			j := i
			j += len(base.SignsArray)
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
