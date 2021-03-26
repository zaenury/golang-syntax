package services

import (
	"context"
	"fmt"
	"log"
	"os"

	model "github.com/zoobc-server-extensions/auto-escrow/models"
	"github.com/zoobc-server-extensions/auto-escrow/util"
	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	rpcModel "github.com/zoobc/zoobc-core/common/model"
	rpcService "github.com/zoobc/zoobc-core/common/service"
)

var (
	general      model.General
	ToBeInserted []mongo.WriteModel
)

//to get block from node address that supplied
func GetBlocksFromNode(chainType int32, limit uint32, height uint32) *rpcModel.GetBlocksResponse {
	nodeaddress := os.Getenv("Node_Address")
	connection := util.GetGrpcConnection(nodeaddress)
	defer connection.Close()

	c := rpcService.NewBlockServiceClient(connection)

	response, err := c.GetBlocks(context.Background(), &rpcModel.GetBlocksRequest{
		ChainType: chainType,
		Limit:     limit,
		Height:    height,
	})

	if err != nil {
		log.Fatalln("Error on inserting new Block", err)
	}

	return response
}

func UpsertBlocks(client *mongo.Client, blocks []model.Block) {
	collection := client.Database("zoosrv_lcl").Collection("Blocks")

	for _, block := range blocks {
		oneBlock := mongo.NewUpdateOneModel()
		oneBlock.SetFilter(bson.M{"ID": block.ID})
		oneBlock.SetUpdate(bson.M{"$set": bson.M{
			"BlockHash":            block.BlockHash,
			"PreviousBlockHash":    block.PreviousBlockHash,
			"Height":               block.Height,
			"Timestamp":            block.Timestamp,
			"BlockSeed":            block.BlockSeed,
			"BlockSignature":       block.BlockSignature,
			"CumulativeDifficulty": block.CumulativeDifficulty,
			"BlocksmithPublicKey":  block.BlocksmithPublicKey,
			"TotalAmount":          block.TotalAmount,
			"TotalFee":             block.TotalFee,
			"TotalCoinBase":        block.TotalCoinBase,
			"Version":              block.Version,
			"PayloadLength":        block.PayloadLength,
			"PayloadHash":          block.PayloadHash,
			"TransactionIDs":       block.TransactionIDs,
		}})
		oneBlock.SetUpsert(true)
		ToBeInserted = append(ToBeInserted, oneBlock)
	}

	// Specify an option to turn the bulk insertion in order of operation
	bulkOption := options.BulkWriteOptions{}
	bulkOption.SetOrdered(true)

	_, err := collection.BulkWrite(context.TODO(), ToBeInserted, &bulkOption)

	if err != nil {
		log.Fatalln("Error on inserting new Block", err)
	}

}

func GetMaxBlockHeight(client *mongo.Client) int {

	collection := client.Database("zoosrv_lcl").Collection("General")

	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&general)
	if err != nil {
		// log.Fatal("blocks error = ", err)
		return 0
	}

	fmt.Println(general)

	return general.LastHeight
}
