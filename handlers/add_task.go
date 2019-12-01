package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"tasks/keeper"
)

func (wrapHandler *WrapperHandler) AddTaskHandler(ctx *fasthttp.RequestCtx) {

	// получение боди из запроса
	reqBody := ctx.Request.Body()
	var reqTask keeper.ReqTask
	err := json.Unmarshal(reqBody, &reqTask)
	if err != nil {
		wrapHandler.Log.Debug().Caller().Err(err).Bytes("тело запроса", reqBody).Msg("")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// сохранение запроса в хранилище
	reqId := wrapHandler.MapTask.SaveTask(keeper.Task{ReqTask: reqTask})

	// создание канала для отслеживания результата обработки задания в воркере
	respStatusChan := make(chan bool)
	defer close(respStatusChan)

	// передача в канал id запроса и канала для ответа
	wrapHandler.ReqExtendedChan <- ReqExtended{ReqId: reqId, RespStatusChan: respStatusChan}

	// получение из канала результата обработки задания в воркере
	respStatus := <-respStatusChan

	if !respStatus {
		wrapHandler.Log.Debug().Caller().Str("address", reqTask.Address).Msg("")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// получение таски из мапы
	task, ok := wrapHandler.MapTask.GetById(reqId)
	if !ok {
		wrapHandler.Log.Debug().Caller().Str("address", reqTask.Address).Msg("")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	// парсинг и формирование ответа
	resp, err := json.Marshal(task.RespTask)
	if err != nil {
		wrapHandler.Log.Debug().Caller().Str("address", reqTask.Address).Msg("")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetBody(resp)
}
