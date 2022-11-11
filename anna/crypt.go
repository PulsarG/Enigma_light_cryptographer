package anna

import (
	"fmt"
	"strconv"
	/* "strings" */
	"test/base"
)

// Строка преобразуется в массив знаков
// Ключевое слово перемешивает массив БазыКонтрольныхЗнаков
// Создается массив порядковых номеров знаков исходного текста в перемешанной БазеКонтрольныхЗнаков
//
//
//
//
//
//

func StartCrypt(text string) string {
	signArr := convertStringToArray(text)
	fmt.Println("Длина текста в символах: ", len(signArr))

	numArr := findNumberSing(signArr)
	fmt.Println(numArr)
	var s string
	for i := 0; i < len(numArr); i++ {
		number := strconv.Itoa(numArr[i])
		s = s + number
	}
	return s
}

func convertStringToArray(s string) []string {
	var signArr []string
	for _, c := range s {
		signArr = append(signArr, fmt.Sprintf("%c", c))
	}
	return signArr
}

func findNumberSing(signArr []string) []int {
	var numArr []int
	newBaseSign := replaceSignArray(5)
	fmt.Println(newBaseSign)
	for numSignText := 0; numSignText < len(signArr); numSignText++ {

		oneSign := signArr[numSignText]

		for i := 0; i < len(newBaseSign); i++ {
			if oneSign == newBaseSign[i] {
				numArr = append(numArr, i)
				break
			}
		}
	}
	return numArr
}

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
