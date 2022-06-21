package raft_node

type CacheRequest struct {
	Key   string `json:"key" binding:"required,min=1"`
	Value string `json:"value" binding:"omitempty,min=1"`
}

func NewCacheRequest() *CacheRequest {
	return &CacheRequest{}
}
