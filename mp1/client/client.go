package main

import (
	"bufio"
	"./mylib/go-shellwords"
	"fmt"
	"log"
	"net/rpc"

	// "flag"
	"os"
	"sync"
	"time"
)

type Log struct {
	Content string
	Lines   int
	Src     string
}

type Query struct {
	Command string
	Args    []string
}

func args_init() *Query {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print(">>")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	}
	cmd := string(sentence[:len(sentence)-1])
	all_in_one, err := shellwords.Parse(cmd)
	fmt.Println(all_in_one)
	//Command line is not valid
	// if len(query) < 4 || err != nil {
	// 	log.Fatal("Invalid Command line")
	// }
	query := &Query{all_in_one[0], all_in_one[1:]}
	return query
}

func remoteQuery(vmID int, query *Query, wg *sync.WaitGroup, acc *[10]int) {
	defer wg.Done()

	remoteAddress := fmt.Sprintf("fa18-cs425-g57-%02d.cs.illinois.edu", vmID)
	// remoteAddress = ""
	// (*query)[3] = fmt.Sprintf("log/vm%d.log", vmID)
	//(*query)[3] = "log/vm1.log"
	// newQuery := make([]string, len(*query))
	// copy(newQuery, *query)
	// newQuery[3] = fmt.Sprintf("log/vm%d.log", vmID)
	client, e := rpc.DialHTTP("tcp", remoteAddress+":12345")

	if e != nil {
		log.Print("Dailling Error: ", e, "\n")
		fmt.Printf("VM %d has failed.\n", vmID)
		return
	}

	reply_log := new(Log)
	// logCall := client.Go("Log.Reply", &newQuery, &reply_log, nil)
	// logCall := client.Go("Log.Reply", &newQuery, &reply_log, nil)
	logCall := client.Go("Log.Reply", query, &reply_log, nil)
	replyCall := <-logCall.Done
	if replyCall == nil {
		log.Print("Request Error: ")
		fmt.Printf("VM %d has failed.\n", vmID)
		return
	}
	acc[vmID-1] = reply_log.Lines
	fmt.Printf("VM# %d Output:\n%s", vmID, reply_log.Content)
	fmt.Printf("VM# %d Total Count: %d\n", vmID, reply_log.Lines)
}

func main() {
	query := args_init()
	var wg sync.WaitGroup
	var accumulator [10]int
	totalAccout := 0
	start := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		// fmt.Println(query)
		go remoteQuery(i, query, &wg, &accumulator)
	}

	wg.Wait()
	end := time.Now()
	latency := end.Sub(start)

	for _, count := range accumulator {
		totalAccout += count
	}
	defer fmt.Printf("Total count of matched pattern: %d.\n", totalAccout)
	defer fmt.Printf("\nLatency of receiving all responses: %d ms.\n", int64(latency)/int64(time.Millisecond))
}

// func QueryTotalCount(pattern string) int {
// 	// query := []string{"grep", "-c", pattern, "log"}
// 	query := Query{"grep", "-c"}
// 	var wg sync.WaitGroup
// 	var accumulator [10]int
// 	totalAccout := 0

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go remoteQuery(i, &query, &wg, &accumulator)
// 	}

// 	wg.Wait()

// 	for _, count := range accumulator {
// 		totalAccout += count
// 	}

// 	return totalAccout
// }
