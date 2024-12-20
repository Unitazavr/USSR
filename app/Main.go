package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите количество столбцов и строк вашей матрицы")
	str, _ := reader.ReadString('\n')
	strslice := strings.Split(str[:len(str)-1], " ")
	if len(strslice) != 2 {
		fmt.Println("Неправильный ввод")
		return
	}
	columns, err := strconv.Atoi(strslice[0])
	rows, err := strconv.Atoi(strslice[1])
	if err != nil {
		fmt.Println("Что то пошло не так")
		return
	}
	var matrix [][]int
	fmt.Println("Вводите свою матрицу, отделяя строки кнопкой enter")
	for i := 0; i < rows; i++ {
		str, _ := reader.ReadString('\n')
		strslice := strings.Fields(str)
		if len(strslice) != columns {
			fmt.Println("Неправильный ввод")
			return
		}
		var row []int
		for j := 0; j < columns; j++ {
			num, _ := strconv.Atoi(strslice[j])
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	// Преобразуем матрицу в JSON
	matrixJSON, err := json.Marshal(matrix)
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON:", err)
		return
	}

	// Отправляем POST-запрос на сервер
	resp, err := http.Post("http://localhost:8087/calculate", "application/json", bytes.NewBuffer(matrixJSON))
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ от сервера
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Сервер вернул ошибку:", resp.Status)
		return
	}

	var result map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println("Ваше среднее арифметическое:", result["average"])
}
