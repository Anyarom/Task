package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
	"tasks/config"
	"tasks/handlers"
	"tasks/keeper"
	"tasks/workers"
)

func main() {

	// настройки логирования в приложении
	log := zerolog.New(os.Stdout).With().Logger()

	// чтение с конфига с помощью библиотеки viper
	cfg, err := config.InitConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("ошибка при чтении конфига")
	}

	// инициализация интерфейса с мапой
	mapKeeper := keeper.InitMapKeeper()

	// создание клиента
	client := http.DefaultClient

	// создание канала для передачи заданий в worker
	reqExtendedChan := make(chan handlers.ReqExtended)
	defer close(reqExtendedChan)

	// запуск в отдельных горутинах нескольких параллельных обработчиков запросов
	for i := 0; i < cfg.Quantity; i++ {
		go workers.Worker(reqExtendedChan, client, mapKeeper, log)
	}

	// инициализация структуры handler для всех запросов
	wrapperHandler := handlers.InitWrapperHandler(log, mapKeeper, client, reqExtendedChan)

	// подключение роутинга к web-серверу
	router := fasthttprouter.New()
	router.POST("/task", handlers.InterceptorLogger(wrapperHandler.AddTaskHandler, log))
	router.GET("/tasks", handlers.InterceptorLogger(wrapperHandler.GetTasksHandler, log))
	router.DELETE("/task", handlers.InterceptorLogger(wrapperHandler.DeleteTaskHandler, log))

	// настройка и запуск сервера
	server := &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: 20 * 1024 * 1024,
	}
	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatal().Caller().Err(err).Msg("Ошибка на сервере")
	}
}
