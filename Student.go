package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}
type University struct { //Структура позволяющая хранить карты студентов
	A map[string]Student
}

func transformation(a string) int { // переводит строчные значения в int
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Второе и третье значение должны быть цифры", err)
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

func (M University) Get(Unit *Student, Name string) { //Метод добавляющий структуру в массив по указателю
	M.A[Name] = *Unit
}
func (M University) Put() {
	fmt.Println("Студенты из хранилища:")
	for _, v := range M.A {
		fmt.Println(v)
	}
}
func main() {

	M := University{ // Массив, принимающий значение
		A: make(map[string]Student),
	}

	sc := bufio.NewScanner(os.Stdin) //Принимает построчный ввод
	for sc.Scan() {
		txt := sc.Text()
		Unit, Name := newStudent(txt)
		M.Get(&Unit, Name) //Вызов метода Get

	}
	M.Put() //Вызов метода Put
}