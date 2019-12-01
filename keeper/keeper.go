package keeper

type ReqTask struct {
	Method  string   `json:"method"`
	Address string   `json:"address"`
	Headers []Header `json:"headers"`
	Body    string   `json:"body"`
}

type RespTask struct {
	ReqId   int      `json:"id"`
	Status  int      `json:"status"`
	Headers []Header `json:"headers"`
	Length  int      `json:"length"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Task struct {
	ReqTask  ReqTask
	RespTask RespTask
}

type Keeper interface {
	SaveTask(task Task) int
	UpdateTask(reqId int, task Task)
	GetById(id int) (task Task, ok bool)
	GetAll() []ReqTask
	DeleteById(id int)
}
