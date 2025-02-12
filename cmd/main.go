package main

import (
	todo "Tudo-list"
	"Tudo-list/pcg/handler"
	"log"
)

func main() {

	//srv := new(todo.Server): создается новый экземпляр структуры Server из пакета todo.
	//if err := srv.Run("8000"); err != nil { ... }:
	//Вызывается метод Run у созданного сервера, передавая порт "8000".
	//	Если метод Run возвращает ошибку (не nil), выполняется блок кода внутри условия.

	//log.Fatal("error running http server: ", err.Error()):
	//Если возникла ошибка при запуске сервера, программа логирует сообщение об ошибке.
	//	log.Fatal выводит сообщение об ошибке и завершает программу с кодом состояния 1.

	//после добавления на sever handler добавляем строчку handlers := new(handler.Handler) и handlers.IntRoutes()
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.IntRoutes()); err != nil {
		log.Fatal("error running http server: ", err.Error())
	}

}

// 2 созданная страница

// https://www.youtube.com/watch?v=Q9hl2oSo0i0&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=2 видео 1   с первой по 8 страницу
