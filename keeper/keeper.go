package keeper

type ReqTask struct {
	Method  string   `json:"method"`
	Address string   `json:"address"`
	Headers []Header `json:"headers"`
	Body    string   `json:"body"`
}

type RespTask struct {
	ReqId   uint64   `json:"id"`
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

type Keeper interface {
	SaveTask(task Task) uint64
	UpdateTask(reqId uint64, task Task)
	GetById(id uint64) (task Task, ok bool)
	GetAll() []ReqTask
	DeleteById(id uint64)
}
