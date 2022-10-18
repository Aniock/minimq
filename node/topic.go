package node

import(
	"fmt"
	"errors"
	// "time"
    "github.com/google/uuid"
)

// var queue = make([]*Message, 100000000, 100000000)

type Topic struct {
	queue []*Message
	name string
	head uint64
	tail uint64
	count uint64
	maxsize uint64
}

func NewTopic(topicName string) * Topic{
	return &Topic{
		queue:make([]*Message, 10, 10),
		name:topicName,
		head:0,
		tail:0,
		count:0,
		maxsize:10,
	}
}

func (t *Topic) GenerateID() string {
	return uuid.New().String()
}

func (t *Topic) PutMessage(m *Message) {
	t.put(m)
	return
}

func (t *Topic) full() bool {
	return t.count == t.maxsize
}

func (t *Topic) empty() bool {
	return t.count == 0
}


func (t *Topic) put(m *Message) bool {
	//判断队列是否已满
	if t.full() == true {
		fmt.Println("满了")
		return false
	}

	//插入队列
	t.queue[t.tail] = m

	t.tail++

	//到达切片尾部绕回头
	if t.tail == t.maxsize {
		t.tail = 0
	}

	t.count++

	return true
}

func (t *Topic) get() (*Message, error) {
	var message *Message

	if t.empty() == true {
		return message, errors.New("null queue")
	}

	message = t.queue[t.head]
	t.queue[t.head] = nil


	t.head++
	if t.head == t.maxsize {
		t.head = 0
	}

	t.count--
	return message, nil
}