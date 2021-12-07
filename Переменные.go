package main

import (
 "fmt"
 "reflect"
)

func main() {
 // #1 ВАРИАНТ СОЗДАНИЯ ПЕРЕМЕННОЙ:
 var all string = "Hello, my name is Bohdan" // Можно указать и тип, и результат
 var half1 int                               // Можно указать тип, но не указывать результат
 var half2 = "4444"                          // Можно указать выражение, но не уточнять тип

 // Обязательно нужно указать или тип или результат переменной

 // #2 ВАРИАНТ СОЗДАНИЯ ПЕРЕМЕННОЙ:
 fisrtone := "hello"                      // Создание переменной всегда начинается с :=
 fisrtone = "new text"                    // Мы можем изменить уже созданную переменную при помощи знака =
 fmt.Println(all, half1, half2, fisrtone) // Выводи на экран результат

 fmt.Println(reflect.TypeOf(half2))
}