package proto

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc"
)

type greeterServer struct {
}

func NewGreeterServer() GreeterServer {
	return greeterServer{}
}

func (greeterServer) mustEmbedUnimplementedGreeterServer() {}

func (s greeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %s", req.GetName())
	res := HelloResponse{
		Result: result,
	}

	return &res, nil
}

func (s greeterServer) Fibonacci(req *FibonacciRequest, stream grpc.ServerStreamingServer[FibonacciResponse]) error {
	for n := uint32(0); n <= req.N; n++ {
		result := fib(n)
		res := FibonacciResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(time.Second)
	}

	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func (s greeterServer) Average(stream grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error {
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}

	res := AverageResponse{
		Result: sum / count,
	}

	return stream.SendAndClose(&res)
}

func (s greeterServer) Sum(stream grpc.BidiStreamingServer[SumRequest, SumResponse]) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		sum += req.Number
		res := SumResponse{
			Result: sum,
		}

		if err = stream.Send(&res); err != nil {
			return err
		}
	}
	return nil
}
