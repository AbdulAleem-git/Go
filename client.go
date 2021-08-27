package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	conn , _ := net.Dial("tcp","127.0.0.1:8080")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Text to send: ")
	text,_ := reader.ReadString('\n')
	fmt.Println(text)
	if conn != nil{
		fmt.Fprintf(conn,text+"\n")
		msg,_ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message from server :%s",msg)
	}
	
}
func firstMissingPositive(nums []int) int {
    neg := 0 
    for i , _ := range nums{
        if v <= 0 {
            swap(&nums[neg] , &nums[i])
            neg++
        }
    }
    for i, v := range nums[neg:]{
        if v < 0 {
            v = v*-1
        }
        if v <= len(nums[neg:]){
            nums[v-1] = nums[v-1] * -1
        }
    }
    fmt.Println(nums)
    
}
func swap( a *int , b *int) {
    *a ,*b = *b ,*a
}


// 3 4 -1 1 
// 0 1  2 3

// -1 -4 3 -1
//  0 1  2  3

// 7 8 9 11 12 
// 0 1 2 3  4


// 0 ->
// -2-3 -4 8 3  2  3 4  5 6 7
//  0 1  2 3 4  5  6  7 8 9 10
//         1 2  3  4  5 6 7 8


curl -X DELETE 127.0.0.1:7201/api/v1/services/m3db/namespace/default


curl -X POST http://localhost:7201/api/v1/database/create -d '{
	"namespaceName": "default",
	"type": "cluster",
	"numShards": 64,
	"replicationFactor": 3,
	"retentionTime": "240h",
	"blockSize": {
	  "time": "2h"
	},
	"hosts": [
		{
		"id": "m3dbnode-0",
		"zone": "embedded",
		"isolationGroup": "ig-0",
		"weight": 100,
		"address": "m3dbnode-0.m3dbnode",
		"port": 9000
	  },
	  {
		"id": "m3dbnode-1",
		"zone": "embedded",
		"isolationGroup": "ig-1",
		"weight": 100,
		"address": "m3dbnode-1.m3dbnode",
		"port": 9000
	  },
	  {
		"id": "m3dbnode-2",
		"zone": "embedded",
		"isolationGroup": "ig-2",
		"weight": 100,
		"address": "m3dbnode-2.m3dbnode",
		"port": 9000
	  }
	]
  }' | jq .