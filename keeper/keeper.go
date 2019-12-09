package keeper

type ReqTask struct {
	Method  string   `json:"method"`
	Address string   `json:"address"`
	Headers []Header `json:"headers"`
	Body    string   `json:"body"`
}

type RespTask struct {
	Status  int      `json:"status"`
	Headers []Header `json:"headers"`
	Length  int      `json:"length"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// структура для записи значения в мапу
type Task struct {
	ReqTask  ReqTask
	RespTask RespTask
}

// структура для запроса getAll
type ReqTaskExtended struct {
	Id      uint64  `json:"request_id"`
	ReqTask ReqTask `json:"request"`
}

type Keeper interface {
	SaveTask(task Task) uint64
	UpdateTask(reqId uint64, task Task)
	GetById(id uint64) (task Task, ok bool)
	GetAll() []ReqTaskExtended
	DeleteById(id uint64)
}
