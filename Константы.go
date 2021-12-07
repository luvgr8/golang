package main

import "fmt"

func main(){
 // КОНСТАНТА
 const staloeznachenie string = "это значение нельзя изменить" // Значение константы изменить нельзя
 const withouttype = "Без типа" // В константе обязательно нужно указать результат, но тип можно не указывать
 fmt.Println(staloeznachenie, withouttype)
}