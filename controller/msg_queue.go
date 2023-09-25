package controller

import(
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"strconv"
	"encoding/json"
	"msg_quing_system/utility"
	// "sync"

)
//function to handling err
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InvokeMesgQue(id int64) (utility.AjaxResponse, error) {
    response := utility.AjaxResponse{Status: "success", Message: "Queue created successfully", Payload: []interface{}{}}
    fmt.Println("Rabbitmq in Golang: Getting started")

    connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Println(err)
        response.Status = "failure"
        response.Message = "Something Went Wrong"
        return response, err
    }
    defer connection.Close()

    fmt.Println("Great success for connection to rabbitmq instance")

    err=Producer(connection, id)
    if err!=nil{
	 	log.Println(err)
        response.Status = "failure"
        response.Message = "Something Went Wrong"
        return response, err
	}
	err=Consumer(connection) // Wait for Consumer to complete
	if err!=nil{
		log.Println(err)
        response.Status = "failure"
        response.Message = "Something Went Wrong"
        return response, err
	}
    return response, err
}


// creating sender to rabbitmq
func Producer(connection *amqp.Connection, id int64)error{
	ch,err:=connection.Channel()//creating a channel
	FailOnError(err,"Failed to open a channel")
	defer ch.Close()

	queueName :="myqueue"
	_,err=ch.QueueDeclare(
		queueName, //quename
		true, //durable
		false, //delete when unused
		false, //exclusive
		false, //no-wait
		nil,   //arguments 
	)
	FailOnError(err,"Failed to declare queue")

	message:= "Hello, RabbitMq! first meetings are always great :)" // msg to send
	msg := fmt.Sprintf("%d - %s", int64(id), message)
	err = ch.Publish( 	// publish msg to the queue
		"",   			//exchange
		queueName, 		//routing key
		false,   		// madatorymsg
		false,  		//immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(msg),
		})
	FailOnError(err,"Failed to publish a message")
	fmt.Printf("send: %s",msg)	
	return err
}

func Consumer(connection *amqp.Connection) error{
	ch, err := connection.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err)
		
	}
	defer ch.Close()

	queueName := "myqueue"
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to declare queue:", err)
		return err
	}

	receiveMsg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to consume messages:", err)
		return err
	}

	go func() {
		for msg := range receiveMsg {
			id, message := ExtractIdAndMsgFromBody(msg.Body)
			fmt.Printf("Received Msg with ID: %d - %s\n", id, message)

			if id != 0 {
				res, err := Db.product.GetProductImagesByID(id)
				if err != nil {
					log.Println("Error fetching images:", err)
				} else if res != nil {
					arrImg, err := BytesToStrings(res)
					if err!=nil{
						log.Println("error convertin",err)
					}else{
					compressedPath, err := CompressAndStoreInLocal(arrImg)
					if err != nil {
						log.Println("Error compressing and storing:", err)
					} else {
						err := Db.product.AddCompressedImagePath(compressedPath, id)
						if err != nil {
							log.Println("Error adding compressed image path:", err)
						}
					}
				}
			}
			}
		}
	}()

	log.Println("Waiting for messages...")
	return err
}


func ExtractIdAndMsgFromBody(body []byte) (int64, string) {
    parts := strings.SplitN(string(body), " - ", 2)
    if len(parts) != 2 {
     FailOnError(nil,"id not received by consumer")
        return -1, ""
    }
    id, err := strconv.ParseInt(parts[0], 10, 64)
    if err != nil {
      FailOnError(err,"parse int err")
        return -1, ""
    }
    
    return id, parts[1]
}


// Convert []byte to []string
func BytesToStrings(data []byte) ([]string,error) {
	var stringSlice []string
	err := json.Unmarshal([]byte(data), &stringSlice)
	if err != nil {
		log.Println(err)
	}
	return stringSlice,err
}
