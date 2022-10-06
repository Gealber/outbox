package cat

import (
	"context"
	"net/http"

	contModel "github.com/Gealber/outbox/controllers/model"
	"github.com/Gealber/outbox/repositories/model"
	"github.com/gin-gonic/gin"
)

type catRepository interface {
	Create(
		ctx context.Context,
		cat model.Cat,
	) (*model.Cat, error)
	Update(
		ctx context.Context,
		id string,
		cat model.Cat,
	) (*model.Cat, error)
}

type Controller struct {
	catRepository catRepository
}

// New ...
func New(catRepository catRepository) *Controller {
	return &Controller{
		catRepository: catRepository,
	}
}

func (ctrl *Controller) Create(c *gin.Context) {
	request := contModel.CreateCatRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

		return
	}

	output, err := ctrl.catRepository.Create(c, request.ToRepoModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, contModel.CatResponse{}.FromRepoModel(output))
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id param not provided in path"})

		return
	}

	request := contModel.UpdateCatRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

		return
	}

	output, err := ctrl.catRepository.Update(c, id, request.ToRepoModel())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, contModel.CatResponse{}.FromRepoModel(output))
}
