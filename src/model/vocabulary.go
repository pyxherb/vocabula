package model

import (
	"context"
	"time"
	"vocabula/common"

	"go.mongodb.org/mongo-driver/bson"
)

type WordRef struct {
	Language string `json:"language" bson:"language"`
	Name     string `json:"name" bson:"name"`
}

type Origin struct {
	Word        WordRef `json:"word" bson:"word"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
}

type Meaning struct {
	Meaning         string   `json:"meaning" bson:"meaning"`
	Origins         []Origin `json:"origins" bson:"origins"`
	LexicalCategory string   `json:"lexicalCategory" bson:"lexical_category"`
}

type Word struct {
	Name     string    `json:"name" bson:"name"`
	Meanings []Meaning `json:"meanings,omitempty" bson:"meanings"`
	Language string    `json:"language" bson:"language"`
}

func QueryWord(language string, name string, wordOut *Word) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"language": language, "name": name}

	err := common.DbConn.VocabularyCollection.FindOne(ctx, filter).Decode(wordOut)
	if err != nil {
		return err
	}

	return nil
}

func UpdateWord(language string, name string, word *Word) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"language": language, "name": name}

	bytes, err := bson.Marshal(word)
	if err != nil {
		return err
	}

	var doc bson.M
	bson.Unmarshal(bytes, &doc)

	_, err = common.DbConn.VocabularyCollection.UpdateOne(ctx, filter, doc)
	if err != nil {
		return err
	}

	return nil
}
