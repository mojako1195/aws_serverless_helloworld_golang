package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type MyEvent struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
}

type MyResponse struct {
    StatusCode string `json:"200:"`
}

type widget struct {
	ID   string `dynamo:"ID"`
    Time time.Time
}
 
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	// DBと接続するセッションを作る→DB接続
    sess := session.Must(session.NewSession())
    db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	table := db.Table("HelloWorldDatabase")

	w := widget{ID: event.FirstName+event.LastName, Time: time.Now()}
	err := table.Put(w).Run()
    if err != nil {
        panic(err)
    }
    return MyResponse{StatusCode: fmt.Sprintf("Hello from Lambda,%s %s", event.FirstName, event.LastName)}, nil
}
 
func main() {
    lambda.Start(HandleLambdaEvent)
}