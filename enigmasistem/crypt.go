package enigmasistem

import (
	"fmt"
	/* "strconv" */ /* "enigma/base" */)

func StartCrypt(text, key string) (string, bool) {
	signArr := convertStringToArray(text)
	key1, key2, key3 := getKeyWordData(key)
	Rotor1, Rotor2, Rotor3 := replacing(key1, key2, key3)

	// Проверка на недопустимые символы
	var s string
	isWrong, wrongSing := checkWrongSing(signArr)
	if !isWrong {
		s = crypting(signArr, Rotor1, Rotor2, Rotor3)
	} else {
		return wrongSing, false
	}

	/* fmt.Println(len(Rotor1), len(Rotor2), len(Rotor3), len(Mirror)) */

	return s, true
}

// Строка преобразуется в массив знаков
func convertStringToArray(s string) []string {
	var signArr []string
	for _, c := range s {
		signArr = append(signArr, fmt.Sprintf("%c", c))
	}
	return signArr
}

// Из ключа получаем три инта
func getKeyWordData(key string) (int, int, int) {
	var x, y, z int

	keyArr := convertStringToArray(key)

	x, y = findNumberFromKey(keyArr)
	z = len(keyArr)

	return x, y, z
}

func findNumberFromKey(arr []string) (int, int) {
	firstLileral := arr[0]
	lastLiteral := arr[(len(arr) - 1)]

	var numFirstLiteral, numLastLiteral int

	for i := 0; i < len(SignsArray); i++ {
		if SignsArray[i] == firstLileral {
			numFirstLiteral = i
		} else {
			continue
		}
	}

	for i := 0; i < len(SignsArray); i++ {
		if SignsArray[i] == lastLiteral {
			numLastLiteral = i
		} else {
			continue
		}
	}

	return numFirstLiteral, numLastLiteral
}

// Перемешивание:
func replacing(key1, key2, key3 int) ([52]string, [52]int, [52]int) {
	// Перемешиваем первым интом Базовый набор - ротор1
	newBaseSign := replaceSignArray(key1, SignsArray)
	// Перемешиваем вторым интом второй Ротор2
	newRotor2 := replaceRotor(key2, Mirror)
	// Перемешиваем третьим интом третий Ротор3
	newRotor3 := replaceRotor(key3, Mirror)

	return newBaseSign, newRotor2, newRotor3
}

// Для каждого знака из массива строки:
func crypting(signArr []string, Rotor1 [52]string, Rotor2, Rotor3 [52]int) string {
	var allResult string
	for i := 0; i < len(signArr); i++ {
		// Поиск Первого Индекса по знаку = х
		indexForRotor2 := findIndexInFirstRotor(signArr[i], Rotor1)
		// Поиск Второго Индекса по значению х = у
		indexForRotor3 := findIndexInRotor(indexForRotor2, Rotor2)
		// Поиск Третьего Индекса по значению у = z
		indexForMirror := findIndexInRotor(indexForRotor3, Rotor3)
		// Поиск Индекса в Отражателе по z = Хх
		indexFromMirror := Mirror[indexForMirror]
		// Поиск значения в Роторе3 2 1 по индексу Хх = Уу
		result := Rotor1[Rotor2[Rotor3[indexFromMirror]]]
		allResult += result
	}
	return allResult
}

func findIndexInFirstRotor(s string, rotor [52]string) int {
	var index int
	for i := 0; i < len(rotor); i++ {
		if rotor[i] == s {
			index = i
			break
		} else {
			continue
		}
	}
	return index
}

func findIndexInRotor(s int, rotor [52]int) int {
	var index int
	for i := 0; i < len(rotor); i++ {
		if rotor[i] == s {
			index = i
			break
		} else {
			continue
		}
	}
	return index
}

func findIndexInMirror(s int, rotor [52]int) int {
	var index int
	for i := 0; i < len(rotor); i++ {
		if rotor[i] == s {
			index = i
			break
		} else {
			continue
		}
	}
	fmt.Println(index)
	return index
}

func CheckLenKey(key string) bool {
	keyArr := convertStringToArray(key)

	if len(keyArr) > len(SignsArray) {
		return false
	} else {
		return true
	}
}

func checkWrongSing(arr []string) (bool, string) {
	check := false
	for indexOfEnterText := 0; indexOfEnterText < len(arr); indexOfEnterText++ {
		for indexOfBaseSingarr := 0; indexOfBaseSingarr < len(SignsArray); indexOfBaseSingarr++ {
			if arr[indexOfEnterText] != SignsArray[indexOfBaseSingarr] {
				check = true
			} else {
				check = false
				break
			}
		}
		if check {
			return check, arr[indexOfEnterText]
		}
	}
	return check, ""
}
