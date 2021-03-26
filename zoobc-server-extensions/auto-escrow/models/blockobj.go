package models

type (
	Block struct {
		ID                   int64    `bson:"ID" json:"ID,omitempty"`
		BlockHash            []byte   `bson:"BlockHash" json:"BlockHash,omitempty"`
		PreviousBlockHash    []byte   `bson:"PreviousBlockHash" json:"PreviousBlockHash,omitempty"`
		Height               uint32   `bson:"Height" json:"Height,omitempty"`
		Timestamp            int64    `bson:"Timestamp" json:"Timestamp,omitempty"`
		BlockSeed            []byte   `bson:"BlockSeed" json:"BlockSeed,omitempty"`
		BlockSignature       []byte   `bson:"BlockSignature" json:"BlockSignature,omitempty"`
		CumulativeDifficulty string   `bson:"CumulativeDifficulty" json:"CumulativeDifficulty,omitempty"`
		BlocksmithPublicKey  []byte   `bson:"BlocksmithPublicKey" json:"BlocksmithPublicKey,omitempty"`
		TotalAmount          int64    `bson:"TotalAmount" json:"TotalAmount,omitempty"`
		TotalFee             int64    `bson:"TotalFee" json:"TotalFee,omitempty"`
		TotalCoinBase        int64    `bson:"TotalCoinBase" json:"TotalCoinBase,omitempty"`
		Version              uint32   `bson:"Version" json:"Version,omitempty"`
		PayloadLength        uint32   `bson:"PayloadLength" json:"PayloadLength,omitempty"`
		PayloadHash          []byte   `bson:"PayloadHash" json:"PayloadHash,omitempty"`
		TransactionIDs       []int64  `bson:"TransactionIDs" json:"TransactionIDs,omitempty"`
		XXX_NoUnkeyedLiteral struct{} `bson:"-" json:"-"`
		XXX_unrecognized     []byte   `bson:"-" json:"-"`
		XXX_sizecache        int32    `bson:"-" json:"-"`
	}

	General struct {
		Key        string `bson:"Key" json:"Key"`
		LastHeight int    `bson:"LastHeight" json:"LastHeight"`
	}
)
