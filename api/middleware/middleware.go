package middleware

import (
	"bufio"
	"github.com/kataras/iris"
	"github.com/timtadh/data-structures/hashtable"
	"github.com/timtadh/data-structures/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type QueueNode struct {
	Domain   string
	Weight   int
	Priority int
}

type PrioNode struct {
	Weight   int
	Priority int
}

var Queue []*QueueNode
var PriorityConfig hashtable.Hash
var priorityConfigInitialSize int

func Init() {
	priorityConfigInitialSize = 10
	PriorityConfig = readPriorityConfig()
}

func Handler(c iris.Context) {
	var domain = c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}

	if !PriorityConfig.Has(types.String(domain)) {
		c.JSON(iris.Map{"status": 400, "result": "unexpected domain"})
		return
	}

	var node = QueueNode{}

	var hashNode, err = PriorityConfig.Get(types.String(domain))

	if err != nil {
		log.Fatal(err)
		c.JSON(iris.Map{"status": 500, "result": "something went wrong"})
		return
	}

	node.Domain = domain
	node.Weight = hashNode.(PrioNode).Weight
	node.Priority = hashNode.(PrioNode).Priority

	Queue = append(Queue, &node)

	sortQueue()

	c.Next()
}

func sortQueue() {
	sort.Slice(Queue, func(i, j int) bool {
		return Queue[i].Weight*Queue[i].Priority > Queue[j].Weight*Queue[j].Priority
	})
}

func readPriorityConfig() hashtable.Hash {
	var path, _ = filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	var scanner = bufio.NewScanner(file)

	var hash = *hashtable.NewHashTable(priorityConfigInitialSize)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		var key = types.String(scanner.Text())
		var node = PrioNode{}
		scanner.Scan()
		node.Weight, _ = strconv.Atoi(strings.Split(scanner.Text(), ":")[1])
		scanner.Scan()
		node.Priority, _ = strconv.Atoi(strings.Split(scanner.Text(), ":")[1])

		hash.Put(key, node)
	}

	return hash
}
