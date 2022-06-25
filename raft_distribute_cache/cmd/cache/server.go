package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/cache"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/config"
	badgerCache "github.com/penk110/interview_ext_go/raft_distribute_cache/internal/cache/badger_cache"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/raft_node"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/middleware"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "c", "../conf/config.default.yaml", "yaml config file ")
	flag.Parse()

	log.Printf("config path: %s\n", configPath)
	conf := config.LoadConfig(configPath)
	// set gin mode
	gin.SetMode(conf.Service.Mode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	rootLogger := conf.Log.NewLogger(strings.ToLower(conf.Service.Mode) == "debug")

	// cache
	badgerCacheClient := badgerCache.NewBadgerCache(conf.RaftConfig.LocalCache)
	// raft node
	raftNode, err := raft_node.NewRaftNode(conf.Service.ServiceID, conf.Service.ServiceName, conf.RaftConfig, badgerCacheClient, conf.Nodes)
	if err != nil {
		log.Fatal(err)
	}

	// 中间件 使得所有请求都被代理到leader
	r.Use(middleware.CacheMiddleware(conf.RaftConfig.Transport, raftNode, conf.Nodes))

	// service
	cacheService := cache.NewCacheService(rootLogger, badgerCacheClient, raftNode)

	// register controller here
	{
		cacherController := cache.NewCacherController(cacheService)
		cacherController.Register(r)
	}

	log.Println("start server...")
	listenAddr := fmt.Sprintf("%s:%s", conf.Service.IP, conf.Service.Port)
	log.Fatal(r.Run(listenAddr))
}
