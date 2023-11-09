package services

import (
	"context"

	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/db"
	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/payload"
)

type IPostService interface {
	CreateNewPost(ctx context.Context, post *payload.PostData) error
}

type post struct {
	dbfactory db.DBFactory
}

func NewPostService(dbfactory db.DBFactory) IPostService {
	return &post{
		dbfactory: dbfactory,
	}
}

func (d *post) CreateNewPost(ctx context.Context, post *payload.PostData) error {
	err := d.dbfactory.GetPostDB().SavePost(ctx, post)
	if err != nil {
		return err
	}
	return nil
}
