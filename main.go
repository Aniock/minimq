package main

import(
	"fmt"
    "strconv"
    "net/http"
    "github.com/gin-gonic/gin"
    "runtime/debug"
)



type QueueMessage struct {
	MessageId int
	Payload string
}

var queue = make([]QueueMessage, 100000000, 100000000)

func main(){

    // var slice []int
    // var queue = make([]QueueMessage, 100000000, 100000000)

    
    fmt.Println( debug.freeOSMemory() )

    for i := 0; i < 100000000; i++ {
        queue[i] = QueueMessage{
            MessageId : i+1,
            Payload   : "abcd",
        }
    }

    // 

    // 初始化 Gin 实例
    r := gin.Default()

    r.GET("/pub", pub_topic)

    r.GET("/get", get_topic)

    // 运行服务
    r.Run("127.0.0.1:8080")
}

//生产消息
func pub_topic(c *gin.Context){
	

    c.JSON(http.StatusOK, gin.H{
        "code"    : 10000,
        "message" : "pub_topic",
        "queue" : queue,
    })
}

//消费消息
func get_topic(c *gin.Context){

    var id = c.DefaultQuery("id", "4399")

    message_id,_ := strconv.Atoi(id);
    queue_message := queue[message_id]

    c.JSON(http.StatusOK, gin.H{
        "code"    : 10000,
        "message" : "get_topic",
        "queue"   : queue_message,
    })
}