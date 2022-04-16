package channel_recive_result

import "log"

type ILooper interface {
	loop()
	cleanup()
	Close()
	Push(j IJob)
}

const (
	_defaultJobChSize = 2048
)

var (
	_defaultLooper ILooper
)

func InitDefaultLooper() ILooper {
	_defaultLooper = &Looper{
		jobCh:  make(chan IJob, _defaultJobChSize),
		doneCh: make(chan struct{}),
	}
	go _defaultLooper.loop()

	return _defaultLooper
}

func NewLooper() {
	return
}

type Looper struct {
	jobCh  chan IJob
	doneCh chan struct{}
}

func (l *Looper) Push(j IJob) {
	l.jobCh <- j
}

func (l *Looper) loop() {
	for {
		select {
		case <-l.doneCh:
			l.cleanup()
			return
		case job := <-l.jobCh:
			go func() {
				job.Handler()
				if err := job.GetErr(); err != nil {
					log.Printf("job: %s handle failed, err: %v", job.GetJobID(), err)
				}
			}()
		}
	}
}

func (l *Looper) Close() {
	close(l.doneCh)
}

func (l *Looper) cleanup() {
	log.Printf("start cleanup")
	for job := range l.jobCh {
		job.Handler()
		if err := job.GetErr(); err != nil {
			log.Printf("job: %s handle failed, err: %v", job.GetJobID(), err)
		}
	}
	log.Printf("cleanup done")
}
