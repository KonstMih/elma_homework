//	Название:
//		"Циклическая ротация"
//
//	Назначение:
//		Сдвиг элементов массива вправо (при этом последний элемент перемещается в начало).
//
//	Описание:
//		Требуемый массив формируется конкатенацией (соединением) двух срезов переданного массива.
//		Первый срез является последним элементом переданного массива. Второй срез является
//		переданным массивом без последнего элемента. В результате получается массив с единичным сдвигом.
//		Далее операция конкатенации повторяется необходимое количество раз.
//
//		Функция возвращает полученный в результате всех сдвигов массив.

package rotation

func Solution(A *[]int) []int {

	var rot []int //	результирующий массив

	//	соединение последнего элемента массива с оставшейся частью необходимое количество раз
	for i := 1; i <= (*A)[0]; i++ {
		rot = append((*A)[(len(*A)-1):], (*A)[:len(*A)-1]...)

	}

	// возвращение полученного результата
	return rot

}
