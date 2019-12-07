package handlers

import (
	"github.com/valyala/fasthttp"
	"strconv"
)

func (wrapHandler *WrapperHandler) DeleteTaskHandler(ctx *fasthttp.RequestCtx) {

	// проверка, есть ли ключ в запросе
	if !ctx.QueryArgs().Has("id") {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// получение ключа из запроса
	idArg := ctx.QueryArgs().Peek("id")[:]
	id, err := strconv.ParseUint(string(idArg), 10, 64)
	if err != nil {
		wrapHandler.Log.Debug().Caller().Err(err).Bytes("id", idArg).Msg("")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	wrapHandler.MapTask.DeleteById(id)
}
