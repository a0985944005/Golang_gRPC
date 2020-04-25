package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"../pb"
)

func main() {
	// 連線到遠端 gRPC 伺服器。
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	// 建立新的 Calculator 客戶端，所以等一下就能夠使用 Calculator 的所有方法。
	c := pb.NewCalculatorClient(conn)

	log.Printf("%T", c)

	plus(c)
	desc(c)
	echo(c)

}
func plus(c pb.CalculatorClient) {
	// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Plus 函式，讓兩個數字相加。
	r, err := c.Plus(context.Background(), &pb.CalcRequest{NumberA: 32, NumberB: 32})

	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}

	log.Printf("Plus回傳結果：%d", r.Result)
}

func desc(c pb.CalculatorClient) {
	// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Desc 函式，讓兩個數字相減。
	r, err := c.Desc(context.Background(), &pb.CalcRequest{NumberA: 64, NumberB: 32})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("Desc回傳結果：%d", r.Result)
}

func echo(c pb.CalculatorClient) {
	// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Desc 函式，讓兩個數字相減。
	r, err := c.Echo(context.Background(), &pb.EchoRequest{Message: "Hello world!"})
	if err != nil {
		log.Fatalf("無法執行 Echo 函式：%v", err)
	}
	log.Println("Echo回傳結果\nMessage：", r.Message, "\nUnixtime", r.Unixtime)
}
