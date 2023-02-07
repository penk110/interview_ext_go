package lib_mongoDB

import (
	"encoding/hex"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"

	"context"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"math/rand"
)

var cli *mongo.Client

func TestInsertMany(t *testing.T) {

}

func init() {
	var err error
	// "mongodb://user:password@localhost:27017".
	//uri := "mongodb://admin:admin@127.0.0.1:27017"
	uri := "mongodb://127.0.0.1:27017"
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// mongodb+srv://<username>:<password>@cluster0-zzart.mongodb.net/test?retryWrites=true&w=majority
	cli, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return
}

func NewDB(db string) *mongo.Database {
	return cli.Database(db)
}

type APIRecord struct {
	TaskID      string        `json:"task_id"`
	URL         string        `json:"url"`
	StatusCode  int           `json:"status_code"`
	StartTS     int64         `json:"start_ts"`
	EndTS       int64         `json:"end_ts"`
	Duration    time.Duration `json:"duration"`
	RespTraffic int64         `json:"resp_traffic"` // 流量统计 header + body
	ReqTraffic  int64         `json:"req_traffic"`
	Body        interface{}   `json:"body"` // TODO: 是否必须保存
}

func TestAsyncResultCollectorPull(t *testing.T) {
	db := NewDB("cobra")

	records := []interface{}{}
	round := 0
	id, _ := uuid.NewV4()
	uid := hex.EncodeToString(id[:])
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	start := time.Now()
	exit := true
	for exit {
		select {
		case <-ticker.C:
			exit = false
			break
		default:
			record := &APIRecord{
				TaskID:      uid,
				URL:         uid,
				StatusCode:  200,
				StartTS:     time.Now().Unix(),
				EndTS:       time.Now().Unix(),
				Duration:    time.Duration(rand.Intn(100)),
				RespTraffic: rand.Int63n(100),
				ReqTraffic:  rand.Int63n(100),
				Body:        `{"code":200,"message":"ok","data":"ok"}`,
			}
			records = append(records, record)
			round++
			if round == 10000 {
				ctx, cancel := context.WithTimeout(context.TODO(), time.Second*60)
				db.Collection("step_records").InsertMany(ctx, records)
				cancel()
				records = []interface{}{}
				round = 0
			}
		}
	}
	end := time.Now()

	t.Logf("cost: %v\n", end.Sub(start))
}

func TestCreateTaskCollection(t *testing.T) {
	db := NewDB("cobra")

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*180)
	defer cancel()

	for i := 0; i < 3; i++ {
		id, _ := uuid.NewV4()
		uid := hex.EncodeToString(id[:])
		err := db.CreateCollection(ctx, "step_records_"+uid)
		if err != nil {
			panic(err)
		}
		t.Log(uid)
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	collections, err := db.ListCollectionNames(ctx, "")
	if err != nil {
		log.Fatal(err)
	}

	t.Log(collections)

}
