package main

import "fmt"

func main() {
 newstring := "first string \n second string \n third string" //  \n используется для перехода на новую строку
 karetka := "first string \r second string \r third string" //    \r: возврат каретки
 tabylacia := "first string \t second string \t third string" //  \t: табуляция
 doublekavichka := "first string \" second string \" third string"  //  \" двойная кавычка внутри строк
 fmt.Println(newstring, karetka, tabylacia,doublekavichka)
}