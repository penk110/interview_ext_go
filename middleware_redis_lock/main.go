package main

import (
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/penk110/interview_ext_go/middleware_redis_lock/lock"
)

var ORM *gorm.DB

func InitDB() error {
	var err error
	dsn := "dev:Dev@3306@tcp(192.168.3.11:3306)/middleware?charset=utf8mb4&parseTime=True&loc=Local"
	ORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Run() {
	cronJob("cronJob_1")
	cronJob("cronJob_2")
}

func job() {
	jobName := "job_" + strconv.FormatInt(time.Now().UnixMicro(), 10)
	locker, err := lock.NewLockerWithAttr(jobName, "/locker/product/pro_1655042378", lock.WhitTTL(time.Second*3), lock.WhitLuaScript(lock.IncrLuaScript)).Lock()
	if err != nil || locker == nil {
		log.Printf("job() jobName: %s, new locker failed, err: %v\n", jobName, err)
		return
	}
	defer locker.Unlock()
	err = ORM.Exec(`update t_products set price=price+1 where pid="pro_1655042378"`).Error
	if err != nil {
		log.Printf("job() update failed, err: %v\n", err)
		return
	}

	log.Printf("job() exec success, jobName: %s\n", jobName)
}

func cronJob(jobName string) {
	c := cron.New(cron.WithSeconds())
	id, err := c.AddFunc("0/5 * * * * *", job)
	if err != nil {
		log.Printf("cronJob() add cronJob success, jobName: %s\n", jobName)
	}
	log.Printf("cronJob() start cronJob success, jobName: %s, jobId: %v\n", jobName, id)
	c.Start()
}

func main() {
	if err := InitDB(); err != nil {
		log.Fatalln(err.Error())
	}

	//go Run()

	var i = 1
	for i <= 100 {
		i++
		go job()
	}

	//go job()
	//go job()

	select {}
}
