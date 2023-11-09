package db

import (
	"context"
	"log"
	"time"

	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/configs"
	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/database"
	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/payload"
)

type IPost interface {
	SavePost(ctx context.Context, post *payload.PostData) error
}

type post struct {
	collection database.DatabaseCollection
}

func NewPostDb(dbclient database.DBClient) IPost {
	return &post{
		collection: dbclient.Collection(configs.POST_COLLECTION),
	}
}

func (p *post) SavePost(ctx context.Context, post *payload.PostData) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	postdata := Post{
		Title:       post.Title,
		Description: post.Description,
		Writer:      post.Writer,
	}

	doc, err := p.collection.InsertOne(ctx, postdata)
	if err != nil {
		return err
	}
	log.Printf("%v", doc)
	return nil
}
