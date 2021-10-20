//	Название:
//		"Домашнее задание".
//
//	Назначение:
//		Решение задач, заданных на курсе.
//
//	Описание:
//		Программа запрашивает название задачи, которую необходимо решить. Затем она отправляет запрос,
//		чтобы получить данные для обработки, и в зависимости от выбранной задачи выводит ответ.

package main

import (
	"bufio"
	"fmt"
	"homework/entry"
	"homework/lost"
	"homework/rotation"
	"homework/sequence"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	taskName := choice()

	var chanData chan []int = make(chan []int)

	go requestData(taskName, chanData)

	fmt.Println(calculation(taskName, chanData))

}

// Ввод названия задачи для решения
func choice() string {

	fmt.Println("Введите название задачи ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()

}

// Получение данных для обработки
func requestData(name string, c chan []int) {

	resp, err := http.Get("http://116.203.203.76:3000/tasks/" + name) //	Получение данных

	if err != nil {
		fmt.Println("error get")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error read")
	}

	// Формирование массива int(ов) из массива byte(ов)
	var lst []int

	for _, i := range body {
		lst = append(lst, int(i))
	}

	c <- lst

}

// Выбор функции для вычисления ответа
func calculation(name string, c chan []int) interface{} {

	data := <-c

	switch name {

	case "Циклическая ротация":
		task1 := rotation.Solution(&data)
		return task1
	case "Чудные вхождения в массив":
		task2 := entry.Solution(&data)
		return task2
	case "Проверка последовательности":
		task3 := sequence.Solution(&data)
		return task3
	case "Поиск отсутствующего элемента":
		task4 := lost.Solution(&data)
		return task4
	default:
		return "Нет ответа"

	}

}
