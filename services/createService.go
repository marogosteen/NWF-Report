package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"example.com/osunnwf.report/repositories"
)

type CreateService struct {
	Title     string
	BestEpoch int
	BestLoss  float64
	Observed  []float64
	Predicted []float64
}

func (s *CreateService) InsertReport() {
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://marogosteen:uiJK....0000@cluster0.n7tlp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("NNWFReport")
	caseListCollection := db.Collection("CaseList")
	reportCollection := db.Collection("Report")
	record := repositories.LearningCollection{
		Title:     s.Title,
		BestEpoch: s.BestEpoch,
		BestLoss:  s.BestLoss,
		Observed:  s.Observed,
		Predicted: s.Predicted,
	}

	reportResult, err := reportCollection.InsertOne(ctx, record)
	if err != nil {
		fmt.Println(err)
		log.Fatal(nil)
	}

	bson := bson.M{
		"title":    s.Title,
		"recordId": reportResult,
	}

	_, err = caseListCollection.InsertOne(ctx, bson)
	if err != nil {
		log.Fatal(err)
	}
}
