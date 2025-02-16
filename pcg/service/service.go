package service

import "Tudo-list/pcg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

//9 страница
//обявим загатовки интерфейса для наших сущьностей
// интерфейсы мы называем исходя из их доменной зоны, уастка бизнес логики приложения за который они отвечают

// позже создадим структуры сервис которая будет собирать все наши сервисы в 1 месте и обявим функцию контруктор

//после создания 10 страницы...
// сервисы будут обращятся к базе данных поетому мы возвращяемся к сервисаи и обявляем указатель на структуру репозитория (repos *repository.Repository) в качестве аргумента конструктора
// это и есть внедрение зависимостей
