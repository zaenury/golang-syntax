package core

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	con "github.com/zoobc-server-extensions/auto-escrow/connection"
	model "github.com/zoobc-server-extensions/auto-escrow/models"
	"github.com/zoobc-server-extensions/auto-escrow/services"
	"github.com/zoobc-server-extensions/auto-escrow/util"
)

var (
	Block  model.Block
	Blocks []model.Block
)

func StartServer() {
	//Get the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accAddress := os.Getenv("Account_Address")
	seedphrase := os.Getenv("Seed_Phrase")
	nodeaddress := os.Getenv("Node_Address")

	s := fmt.Sprintf("Address = %v & seedPhrase = %v & nodeAddress = %v\n", accAddress, seedphrase, nodeaddress)
	io.WriteString(os.Stdout, s)

	//get the block
	conn := con.Connect()
	height := services.GetMaxBlockHeight(conn)
	fmt.Println("Height = ", height)
	blocks := services.GetBlocksFromNode(0, 100, uint32(height))

	for _, i := range blocks.Blocks {
		block := model.Block{
			ID:                   i.ID,
			BlockHash:            i.BlockHash,
			PreviousBlockHash:    i.PreviousBlockHash,
			Height:               i.Height,
			Timestamp:            i.Timestamp,
			BlockSeed:            i.BlockSeed,
			BlockSignature:       i.BlockSignature,
			CumulativeDifficulty: i.CumulativeDifficulty,
			BlocksmithPublicKey:  i.BlocksmithPublicKey,
			TotalAmount:          i.TotalAmount,
			TotalFee:             i.TotalFee,
			TotalCoinBase:        i.TotalCoinBase,
			Version:              i.Version,
			PayloadLength:        i.PayloadLength,
			PayloadHash:          i.PayloadHash,
			TransactionIDs:       i.TransactionIDs,
		}
		Blocks = append(Blocks, block)
	}

	//set the highest block on general collection
	key := util.LastBlockKey()
	_, max := util.FindMaxAndMinOfSliceOfStruct(Blocks)
	services.UpsertGeneral(conn, key, max.Height)

	//insert the blocks to blocks collection
	services.UpsertBlocks(conn, Blocks)
}
