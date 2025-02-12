package todo

import (
	"context"
	"net/http"
	"time"
)

//Определяется структура Server, которая содержит указатель на http.Server. Это позволяет инкапсулировать стандартный HTTP-сервер Go и добавить к нему дополнительную функциональность.

type Server struct {
	httpServer *http.Server
}

//Метод Run инициализирует и запускает HTTP-сервер:
//Создается новый экземпляр http.Server с следующими настройками:
//Addr: адрес сервера (: + номер порта)
//MaxHeaderBytes: максимальный размер заголовка запроса (1 МБ)
//ReadHeaderTimeout: таймаут чтения заголовка (10 секунд)
//WriteTimeout: таймаут записи ответа (10 секунд)
//Вызывается метод ListenAndServe(), который запускает сервер и начинает прослушивание входящих соединений.

// после завершения 5 страницы добавляем handler http.Handler и Handler: handler,

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20, // 1 MB
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

//Метод Shutdown обеспечивает корректное завершение работы сервера:
//Принимает контекст ctx, который может использоваться для установки таймаута или отмены операции завершения.
//Вызывает метод Shutdown у http.Server, который останавливает прием новых соединений, закрывает все активные соединения и ожидает завершения всех текущих запросов.

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

//  1 созданная страница
