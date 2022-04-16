package channel_recive_result

import (
	"github.com/google/uuid"
	"log"
)

type IJob interface {
	Handler()
	CallBack()
	SetErr(err error)
	GetErr() error
	GetJobID() string
}

func NewJob() IJob {
	return &DefJob{
		jobID: uuid.New().String(),
	}
}

type DefJob struct {
	jobID string
	err   error
}

func (job *DefJob) Handler() {
	var err error
	defer job.CallBack()

	log.Printf("start handle job: %s\n", job.jobID)

	// ... ...
	if err != nil {
		job.SetErr(err)
		return
	}
	return
}

func (job *DefJob) SetErr(err error) {
	job.err = err
}

func (job *DefJob) GetErr() error {
	return job.err
}

func (job *DefJob) CallBack() {
	log.Printf("job: %s callback\n", job.jobID)
	return
}

func (job *DefJob) GetJobID() string {
	return job.jobID
}
