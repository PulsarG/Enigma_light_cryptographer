package anna

import (
	"fmt"
	"strconv"
	"test/base"
)

func StartCrypt(text, key string) (string, bool) {
	signArr := convertStringToArray(text)
	key1, key2, key3 := getKeyWordData(key)
	Rotor1, Rotor2, Rotor3 := replacing(key1, key2, key3)

	s := crypting(signArr, Rotor1, Rotor2, Rotor3)

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

	for i := 0; i < len(arr); i++ {
		if arr[i] == firstLileral {
			numFirstLiteral = i
		} else {
			continue
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == lastLiteral {
			numLastLiteral = i
		} else {
			continue
		}
	}

	return numFirstLiteral, numLastLiteral
}

// Перемешивание:
func replacing(key1, key2, key3 int) ([51]string, [51]string, [51]string) {
	// Перемешиваем первым интом Базовый набор - ротор1
	newBaseSign := replaceSignArray(key1, base.SignsArray)
	// Перемешиваем вторым интом второй Ротор2
	newRotor2 := replaceSignArray(key2, base.Rotor2)
	// Перемешиваем третьим интом третий Ротор3
	newRotor3 := replaceSignArray(key3, base.Rotor3)

	return newBaseSign, newRotor2, newRotor3
}

// Для каждого знака из массива строки:
func crypting(signArr []string, Rotor1, Rotor2, Rotor3 [51]string) string {
	var allResult string
	for i := 0; i < len(signArr); i++ {
		// Поиск Первого Индекса по знаку = х
		indexForRotor2 := findIndexInRotor(signArr[i], Rotor1)
		i2 := strconv.Itoa(indexForRotor2)
		// Поиск Второго Индекса по значению х = у
		indexForRotor3 := findIndexInRotor(i2, Rotor2)
		i3 := strconv.Itoa(indexForRotor3)
		// Поиск Третьего Индекса по значению у = z
		indexForMirror := findIndexInRotor(i3, Rotor3)
		iM := strconv.Itoa(indexForMirror)
		// Поиск Индекса в Отражателе по z = Хх
		indexFromMirror := findIndexInMirror(iM, base.Mirror)
		// Поиск значения в Роторе3 по индексу Хх = Уу
		varFromRotor3 := findVarInRotor(indexFromMirror, Rotor3)
		// Поиск значения в Роторе2 по индексу Уу = Zz
		varFromRotor2 := findVarInRotor(int(varFromRotor3), Rotor2)
		// Получение Результата-Знака в Роторе1 - Базовом наборе по индексу Zz = Result
		result := Rotor1[varFromRotor2]
		allResult += result
	}
	return allResult
}

func findIndexInRotor(s string, rotor [51]string) int {
	var index int
	for i := 0; i < len(rotor); i++ {
		if rotor[i] == s {
			fmt.Println(i, s)
			index = i
			break
		} else {
			continue
		}
	}
	return index
}

func findIndexInMirror(s string, rotor []string) int {
	var index int
	for i := 0; i < len(rotor); i++ {
		if rotor[i] == s {
			fmt.Println(i, s)
			index = i
			break
		} else {
			continue
		}
	}
	return index
}

func findVarInRotor(i int, rotor [51]string) int64 {
	x := rotor[i]
	y, _ := strconv.ParseInt(x, 6, 12)
	return y
}
