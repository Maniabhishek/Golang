package services

import "github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/db"

type Servicefactory interface {
	GetPostService() IPostService
}

type serviceFactory struct {
	postservice IPostService
}

func NewServiceFactory(dbfactory db.DBFactory) Servicefactory {
	return &serviceFactory{
		postservice: NewPostService(dbfactory),
	}
}

func (s *serviceFactory)GetPostService() IPostService {
	return s.postservice
}