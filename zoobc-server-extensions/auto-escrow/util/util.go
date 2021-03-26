package util

import (
	"log"

	model "github.com/zoobc-server-extensions/auto-escrow/models"
	"google.golang.org/grpc"
)

func GetGrpcConnection(address string) *grpc.ClientConn {
	connection, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("grpc connection error: %v", err)
	}
	return connection
}

func FindMaxAndMinOfSliceOfStruct(blocks []model.Block) (min model.Block, max model.Block) {
	min = blocks[0]
	max = blocks[0]
	for _, block := range blocks {
		if block.Height > max.Height {
			max = block
		}
		if block.Height < min.Height {
			min = block
		}
	}
	return min, max
}

func LastBlockKey() string {
	return "LastBlockHeight"
}
