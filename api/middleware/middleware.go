package middleware

import (
	"fmt"
	"github.com/kataras/iris"
)

type QueueNode struct {
	Domain   string
	Weight   int
	Priority int
}

var Queue []*QueueNode

type Repository interface {
	Read() []*QueueNode
}

func (q *QueueNode) Read() []*QueueNode {
	return MockQueue()
}

func MockQueue() []*QueueNode {
	return []*QueueNode{
		{
			Domain:   "alpha",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "beta",
			Weight:   5,
			Priority: 5,
		},
		{
			Domain:   "omega",
			Weight:   5,
			Priority: 1,
		},
	}
}

func Init() {
	Queue = append(Queue, &QueueNode{})
}

func Handler(c iris.Context) {
	var domain = c.GetHeader("domain")
	var repository Repository
	repository = &QueueNode{}
	fmt.Println(domain)
	fmt.Println(repository.Read())
	c.Next()
}
