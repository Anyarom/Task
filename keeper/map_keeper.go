package keeper

import (
	"github.com/rs/zerolog/log"
	"sync"
)

type MapKeeper struct {
	lock      sync.Mutex
	currentId uint64
	MapTasks  sync.Map
}

// функция для инициализации структуры MapKeeper
func InitMapKeeper() *MapKeeper {
	var mapTasks sync.Map
	mapKeeper := MapKeeper{MapTasks: mapTasks}
	return &mapKeeper
}

// метод для сохранения в мапу
func (mk *MapKeeper) SaveTask(task Task) uint64 {
	// генерация id инкрементально
	mk.lock.Lock()
	reqId := mk.currentId + 1
	mk.currentId = reqId
	mk.MapTasks.Store(reqId, task)
	mk.lock.Unlock()
	return reqId
}

// метод для обновления значений в мапе
func (mk *MapKeeper) UpdateTask(reqId uint64, task Task) {
	mk.MapTasks.Store(reqId, task)
}

// метод получения одного запроса
func (mk *MapKeeper) GetById(reqId uint64) (task Task, ok bool) {
	t, ok := mk.MapTasks.Load(reqId)
	if !ok {
		log.Debug().Caller().Msg("")
		return Task{}, false
	}
	return t.(Task), true
}

// метод для получения всех запросов из мапы
func (mk *MapKeeper) GetAll() []ReqTaskExtended {
	var tasks []ReqTaskExtended

	mk.MapTasks.Range(func(id, value interface{}) bool {
		reqTaskExtended := ReqTaskExtended{Id: id.(uint64), ReqTask: value.(Task).ReqTask}
		tasks = append(tasks, reqTaskExtended)

		return true
	})
	return tasks
}

// метод для удаления из мапы
func (mk *MapKeeper) DeleteById(reqId uint64) {
	mk.MapTasks.Delete(reqId)
}
