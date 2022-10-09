package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func transformation(a string) int { // переводит строчные значения в int
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Второе и третье значение должны быть цифры")
		log.Fatal(err)
	}
	return b
}

func newStudent(txt string) (Student, string) { // принимает на вход строку, преобразует в структуру
	txtSplited := strings.Split(txt, " ")

	var StudentName string = txtSplited[0]
	return Student{
		txtSplited[0],
		transformation(txtSplited[1]),
		transformation(txtSplited[2])}, StudentName
}

func main() {

	M := make(map[string]*Student) // Массив, принимающий значение по указателю

	sc := bufio.NewScanner(os.Stdin) //Принимает построчный ввод
	for sc.Scan() {
		txt := sc.Text()
		A, studentName := newStudent(txt)
		M[studentName] = &A // Запись в массив через указатель
	}
	fmt.Println("Студенты из хранилища:")
	for _, v := range M {
		fmt.Println(v)
	}
}
