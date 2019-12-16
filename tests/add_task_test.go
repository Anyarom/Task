package tests

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog"
	"io/ioutil"
	"net/http"
	"os"
	"tasks/keeper"
	"testing"
	"time"
)

// создадание интеграционный тест для запроса на создание таски
func TestAddTask(t *testing.T) {

	// настройки логирования
	log := zerolog.New(os.Stdout).With().Logger()

	for i := 0; i < 20; i++ {
		go performRequest(log, t)
	}
	time.Sleep(10 * time.Second)
}

func performRequest(log zerolog.Logger, t *testing.T) {

	// создадание httpClient
	client := http.DefaultClient

	// боди в запросе
	bodyRequest := []byte(`{
			  "method": "POST",
			  "address": "http://google.com",
			  "headers": [
				{
				  "key": "key1",
				  "value": "val1"
				},
				{
				  "key": "key2",
				  "value": "val2"
				}
			  ],
			  "body": "таска"
			}`)

	body := bytes.NewReader(bodyRequest)
	// отправление, используя созданный httpClient, запроса на сервер
	resp, err := client.Post("http://127.0.0.1:8080/task", "application/json", body)
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		t.Fail()
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("")
		}
	}()

	// достать боди ответа
	stat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		t.Fail()
	}

	// парсинг боди ответа
	var responseTask keeper.RespTask
	err = json.Unmarshal(stat, &responseTask)
	if err != nil {
		log.Error().Caller().Err(err).Msg("")
		t.Fail()
	}

	log.Debug().Caller().Msg("success")
}
