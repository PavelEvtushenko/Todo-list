package handler

import (
	"Tudo-list/pcg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	//после изменений на 10 странице
	services *service.Service
}

// создаем функцию после того как заполнили структуру хендлер
func NewHendler(services *service.Service) *Handler {

	return &Handler{services: services}

}

// здесь мы настраивем и инициализируем конечные точки нашего приложения

func (h *Handler) IntRoutes() *gin.Engine {
	//для инициализации роутера вызовим метод гин нью
	router := gin.New()

	//далее обявим наши методы сгрупировав их по маршрутам

	//auth используем для регистрации и авторизации

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/sing-in", h.signIn)
	}

	//будет использоватся для работы со списками и их задачами

	api := router.Group("/api")
	{
		// внутри апи создадим группу лист для работы со списками
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			// создаем группу для задач списков
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}

// 5  страница

//после 8 страницы добавляем h.-----  в методы
// после того как мы создали ендпоинты на 6, 7, 8 странице добаваляем из в качестве аргумена в методы хендлера
