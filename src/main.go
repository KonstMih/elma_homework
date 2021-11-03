//	Название:
//		"Домашнее задание".
//
//	Назначение:
//		Создание микросервиса для решения задач.
//
//	Описание:
//		Сервис принимает запрос в виде названия одной из задач (запрос необходимо написать на кириллице).
//		Затем сервис в зависимости от задачи получает данные, обрабатывает их и посылает в ответ полученный результат.

package main

import (
	"encoding/hex"
	"fmt"
	"homework/entry"
	"homework/lost"
	"homework/rotation"
	"homework/sequence"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", myHandler) //	Обработка запроса

	http.ListenAndServe(":8080", nil) // Запуск сервера

}

// Функция преобразования URL запроса
func hexString(s string) string {

	var str string

	for _, i := range s {
		n := string(i)
		if n != "%" {
			str = str + n
		}
	}

	decoded, _ := hex.DecodeString(str)

	return (string(decoded))
}

// Функция получения данных
func GetData(s string) []int {

	resp, err := http.Get("http://116.203.203.76:3000/tasks/" + s) //	Адрес сервера с данными

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

	return lst
}

// Функция выбора решения для вычисления ответа
func calculation(name string, list []int) []byte {

	var resp []byte

	switch name {

	case "Циклическая ротация":
		task1 := rotation.Solution(&list)
		for _, i := range task1 {
			resp = append(resp, byte(i))
		}
	case "Чудные вхождения в массив":
		task2 := entry.Solution(&list)
		resp = append(resp, byte(task2))
	case "Проверка последовательности":
		task3 := sequence.Solution(&list)
		resp = append(resp, byte(task3))
	case "Поиск отсутствующего элемента":
		task4 := lost.Solution(&list)
		resp = append(resp, byte(task4))
	default:
		resp = []byte("Нет решения")

	}

	return resp
}

// Функция обработки запроса
func myHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.RequestURI()[1:] //	Получение URL запроса

	uStr := hexString(u) //	Преобразование URL в строку (определение задачи)

	data := GetData(uStr) //	Получение данных для задачи

	answer := calculation(uStr, data) // Решение задачи

	fmt.Println(answer)

	w.WriteHeader(http.StatusOK)
	w.Write(answer) //	Отправка ответа

}
