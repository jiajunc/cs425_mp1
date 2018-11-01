package main

import (

	// "errors"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Query struct {
	Command string
	Args    []string
}

type Log struct {
	Content string
	Lines   int
	Src     string
}

func (t *Log) Reply(query Query, reply *Log) error {
	fmt.Printf("%v", query)
	// for _, ele := range query {
	// 	fmt.Println(ele)
	// }

	// home := os.Getenv("/")
	err := os.Chdir("/home/jl46/cs425-mp2/")
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command(query.Command, query.Args...)
	// cmd.Dir = "/home/jl46/"
	var out bytes.Buffer
	cmd.Stdout = &out
	e := cmd.Run()
	if e != nil {
		log.Print("Command Fails ", e)
		return nil
	}
	(*reply).Content = out.String()

	if query.Command == "grep" {
		if query.Args[0] == "-c" {
			outputCount, _ := strconv.Atoi(strings.Trim(reply.Content, "\n"))
			fmt.Print(reply.Content)
			fmt.Print(outputCount)
			(*reply).Lines = outputCount
		} else {
			(*reply).Lines = strings.Count(reply.Content, "\n")
		}
	}
	fmt.Printf("From Command: %d\n", reply.Lines)
	return nil
}

func main() {
	log_reply := new(Log)
	rpc.Register(log_reply)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":12345")
	if e != nil {
		log.Fatal("Listen Error: ", e)
	}
	http.Serve(l, nil) // what is l

}
