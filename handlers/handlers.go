package handlers

import (
	"github.com/rs/zerolog"
	"net/http"
	"tasks/keeper"
)

type WrapperHandler struct {
	Log             zerolog.Logger
	MapTask         keeper.Keeper
	Client          *http.Client
	ReqExtendedChan chan ReqExtended
}

type ReqExtended struct {
	ReqId          uint64
	RespStatusChan chan bool
}

func InitWrapperHandler(log zerolog.Logger, mapTask keeper.Keeper, client *http.Client, reqExtendedChan chan ReqExtended) *WrapperHandler {
	return &WrapperHandler{Log: log, MapTask: mapTask, Client: client, ReqExtendedChan: reqExtendedChan}
}
