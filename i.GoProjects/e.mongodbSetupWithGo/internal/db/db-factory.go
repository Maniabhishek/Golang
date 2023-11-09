package db

import "github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/database"

type DBFactory interface {
	GetPostDB() IPost
}

type dbfactory struct {
	postdb IPost
}

func NewDbFactory(dbclient database.DBClient) DBFactory {
	return &dbfactory{
		postdb: NewPostDb(dbclient),
	}
}

func (d *dbfactory) GetPostDB() IPost {
	return d.postdb
}
