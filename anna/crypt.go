package anna

import (
	"fmt"
	"strconv"
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

func StartCrypt(text, key string) string {
	lenKeyWord := getKeyWordData(key)

	signArr := convertStringToArray(text)
	fmt.Println("Длина текста в символах: ", len(signArr))

	numArr := findNumberSing(signArr, lenKeyWord)
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

func findNumberSing(signArr []string, lenKey int) []int {
	var numArr []int
	newBaseSign := replaceSignArray(lenKey)
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
func getKeyWordData(key string) int {
	lenKeyWord := len(convertStringToArray(key))
	return lenKeyWord
}
