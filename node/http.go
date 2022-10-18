package node

import(
	// "fmt"
    // "log"
    "net/http"
    "io/ioutil"
    "github.com/gin-gonic/gin"
    // "strconv"
    // "github.com/gin-gonic/gin"
    // "runtime/debug"
)

func HttpServer(){
    router := gin.Default()

    router.GET("/ping", ping)
    router.GET("/info", info)
    router.GET("/stats", stats)

    //发布消息
    router.GET("/pub", pub)
    router.GET("/get", get)


    router.NoRoute(NoRoute)

    router.Run("127.0.0.1:4399")
}

func ping(c *gin.Context){
    c.String(http.StatusOK, "ok")
}

//{"version":"1.2.1","broadcast_address":"USER-EUF907CI09","hostname":"USER-EUF907CI09","http_port":4151,"tcp_port":4150,"start_time":1666000769}
func info(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
        "version"    : "1.0.0",
        "start_time" : 1234567890,
    })
}

func stats(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
        "version" : "system统计",
    })
}


/** 
 * http://127.0.0.1:4399/pub?topic=abcd
 */
func pub(c *gin.Context){
    //body大小检查

    //消息空检查

    //读取Topic
    topicName := c.Query("topic")
    if( topicName == "" ){
        c.JSON(http.StatusOK, gin.H{
            "message" : "need query param `topic`",
        })
        return
    }

    //topic初始化
    // topic := NewTopic(topicName)
    topic := GetTopic(topicName)

    //消息初始化
    body,_ := ioutil.ReadAll(c.Request.Body)
    msg := NewMessage(topic.GenerateID(), body)

    topic.PutMessage(msg)

    c.JSON(http.StatusOK, gin.H{
        "version" : "pub消息",
        "topic"   : topicName,
    })
}

func get(c *gin.Context){
    topic := GetTopic("abcd")

    message, _ := topic.get()

    c.JSON(http.StatusOK, gin.H{
        "message"   : message,
    })
    
}



func NoRoute(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
        "message" : "NOT_FOUND",
    })
}