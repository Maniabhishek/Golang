package apis

import (
	"encoding/json"
	"io"
	"log"

	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/payload"
	"github.com/Maniabhishek/Golang/i.GoProjects/serverWithEcho/internal/services"
	"github.com/labstack/echo"
)

type postApi struct {
	postService services.IPostService
}

func NewPostAPI(serviceFactory services.Servicefactory) *postApi {
	return &postApi{
		postService: serviceFactory.GetPostService(),
	}
}

func (p *postApi) SavePost(ctx echo.Context) error {
	// context := context.Background()
	context := ctx.Request().Context()
	body := ctx.Request().Body

	bbytes, error := io.ReadAll(body)
	if error != nil {
		return ctx.JSON(500, "internal server error")
	}

	var postData payload.PostData
	err := json.Unmarshal(bbytes, &postData)
	if err != nil {
		log.Panic(err)
		return ctx.JSON(500, err)
	}
	if err := p.postService.CreateNewPost(context, &postData); err != nil {
		return ctx.JSON(500, err.Error())
	}
	return ctx.JSON(200, map[string]string{
		"message": "post create successfully",
	})
}
