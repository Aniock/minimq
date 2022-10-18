package node

import(
	"fmt"
)

// var topics = make([]*Topic, 10000, 10000)
var topics = make(map[string]*Topic)


//集中管理Topic
func GetTopic(topicName string) *Topic {

	t, ok := topics[topicName]

	fmt.Println(ok)

	if( ok ){
		return t
	}

	//新的Topic
	topic := NewTopic(topicName)
	topics[topicName] = topic

	return topic
}