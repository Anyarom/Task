package workers

import (
	"github.com/rs/zerolog"
	"net/http"
	"strings"
	"tasks/handlers"
	"tasks/keeper"
)

// обработчик запросов
func Worker(reqCh <-chan handlers.ReqExtended, client *http.Client, myMap keeper.Keeper, log zerolog.Logger) {
	for {

		// получение из канала id таски и канала для ответа
		requestWithChannel := <-reqCh

		// получение таски из мапы
		task, ok := myMap.GetById(requestWithChannel.ReqId)
		if !ok {
			log.Debug().Caller().Msg("")
			requestWithChannel.RespStatusChan <- false
			continue
		}

		// выполнение запроса
		resp, success := performRequest(task.ReqTask, requestWithChannel, client, log)
		if !success {
			continue
		}

		// приведение к типу Header для заполнения структуры в ответе
		var hd []keeper.Header
		for key, value := range resp.Header {
			hd = append(hd, keeper.Header{Key: key, Value: value[0]})
		}

		// формирование результата выполнения таски
		respTask := keeper.RespTask{ReqId: requestWithChannel.ReqId, Status: resp.StatusCode, Headers: hd, Length: int(resp.ContentLength)}

		// обновление данных в мапе
		myMap.UpdateTask(requestWithChannel.ReqId, keeper.Task{ReqTask: task.ReqTask, RespTask: respTask})

		// отправление в канал результата об обновлении значения в мапе
		requestWithChannel.RespStatusChan <- true
	}
}

func performRequest(reqTask keeper.ReqTask, requestWithChannel handlers.ReqExtended, client *http.Client, log zerolog.Logger) (*http.Response, bool) {

	// создание нового запроса согласно таски
	req, err := http.NewRequest(reqTask.Method, reqTask.Address, strings.NewReader(reqTask.Body))
	if err != nil {
		requestWithChannel.RespStatusChan <- false
		return nil, false
	}
	headers := map[string][]string{}
	for _, h := range reqTask.Headers {
		headers[h.Key] = []string{h.Value}
	}
	req.Header = headers

	// отправка запроса и получение ответа
	respForClient, err := client.Do(req)
	if err != nil {
		requestWithChannel.RespStatusChan <- false
		return nil, false
	}

	// закрытие боди
	defer func() {
		err := respForClient.Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	}()

	return respForClient, true
}
