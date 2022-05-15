package main

import (
	"github.com/casbin/casbin"
	"log"
)

func main() {
	sub := "manager"
	obj := "/dept"
	act := "POST"

	enForce := casbin.NewEnforcer("model.conf", "permission.csv")

	ok := enForce.Enforce(sub, obj, act)
	if ok {
		log.Println("have access permission")
	}
}
