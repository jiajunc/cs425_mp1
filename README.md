# MP1 - Query

## Environment

Go 1.9.4

## Communication Framework

RPC

## External Library

shellwords - for string parsing

## Usage

```bash
# build
# in server/
go build # generate executeble
# in client/
go build # generate eexecuteble

# run
# in server/
./server
#or without building: 
go run server.go
# in client/
./client
>>grep -nr 000 log/
>>grep -c 000 log/
```
