package middleware

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/raft"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/config"
)

// CacheMiddleware 将请求代理到主节点

func CacheMiddleware(transport string, raftNode *raft.Raft, nodes []*config.Node) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// is leader
		if raftNode.Leader() == raft.ServerAddress(transport) {
			ctx.Next()
			return
		}
		leaderHTTP := GetLeaderHTTP(raftNode, nodes)
		// 肯定是存在leader
		addr, _ := url.Parse(leaderHTTP)
		proxy := httputil.NewSingleHostReverseProxy(addr)
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
		ctx.Abort()
	}
}

func GetLeaderHTTP(raftNode *raft.Raft, nodes []*config.Node) string {
	for _, node := range nodes {
		if string(node.Address) == string(raftNode.Leader()) {
			return node.Http
		}
	}
	return ""
}
