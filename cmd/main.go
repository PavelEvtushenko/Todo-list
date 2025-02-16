package main

import (
	todo "Tudo-list"
	"Tudo-list/pcg/handler"
	"Tudo-list/pcg/repository"
	"Tudo-list/pcg/service"
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
	//после изменения на 10 сранице и хендлерах удаляем  handlers := new(handler.Handler) и добавляем указатели на сервисы
	// обявляем все зависимости в нужном порядке создав репозиторий, сервис который зависит от репозитория а потом хендлер который зависит от сервиса

	//Эта строка создает экземпляр репозитория, который отвечает за взаимодействие с базой данных или другим хранилищем данных.
	repos := repository.NewRepository()
	//Здесь создается экземпляр сервиса, который содержит бизнес-логику приложения. Сервису передается репозиторий, чтобы он мог работать с данными.
	services := service.NewService(repos)
	//Эта строка создает экземпляр обработчика HTTP-запросов. Обработчику передается сервис, чтобы он мог использовать бизнес-логику для обработки запросов.
	handlers := handler.NewHendler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.IntRoutes()); err != nil {
		log.Fatal("error running http server: ", err.Error())
	}

}

// 2 созданная страница

// https://www.youtube.com/watch?v=Q9hl2oSo0i0&list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8&index=2 видео 1   с первой по 8 страницу
// https://www.youtube.com/watch?v=Of7MAgRRGmg видео с 1 по 10 страницу
