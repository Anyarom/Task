package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"tasks/keeper"
)

func (wrapHandler *WrapperHandler) GetTasksHandler(ctx *fasthttp.RequestCtx) {

	// получение всех тасок из мапы
	reqTasks := wrapHandler.MapTask.GetAll()

	// заполнение структуры для ответа
	type RespTasks struct {
		ReqTasks []keeper.ReqTaskExtended `json:"tasks"`
	}
	respTasks := RespTasks{ReqTasks: reqTasks}

	// формирование ответа
	resp, err := json.Marshal(respTasks)
	if err != nil {
		wrapHandler.Log.Debug().Caller().Err(err).Msg("")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetBody(resp)
}
