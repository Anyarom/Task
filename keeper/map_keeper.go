package keeper

import (
	"github.com/rs/zerolog/log"
	"math/rand"
	"sync"
	"time"
)

type MapKeeper struct {
	MapTasks sync.Map
}

// функция для инициализации структуры MapKeeper
func InitMapKeeper() *MapKeeper {
	var mapTasks sync.Map
	mapKeeper := MapKeeper{MapTasks: mapTasks}
	return &mapKeeper
}

// метод для сохранения в мапу
func (mk *MapKeeper) SaveTask(task Task) int {
	// генерация id
	rand.Seed(time.Now().UnixNano())
	reqId := rand.Int()
	mk.MapTasks.Store(reqId, task)
	return reqId
}

// метод для обновления значений в мапе
func (mk *MapKeeper) UpdateTask(reqId int, task Task) {
	mk.MapTasks.Store(reqId, task)
}

// метод получения одного запроса
func (mk *MapKeeper) GetById(reqId int) (task Task, ok bool) {
	t, ok := mk.MapTasks.Load(reqId)
	if !ok {
		log.Debug().Caller().Msg("")
		return Task{}, false
	}
	return t.(Task), true
}

// метод для получения всех запросов из мапы
func (mk *MapKeeper) GetAll() []ReqTask {
	var reqTasks []ReqTask
	mk.MapTasks.Range(func(id, value interface{}) bool {
		reqTasks = append(reqTasks, value.(Task).ReqTask)
		return true
	})
	return reqTasks
}

// метод для удаления из мапы
func (mk *MapKeeper) DeleteById(reqId int) {
	mk.MapTasks.Delete(reqId)
}
