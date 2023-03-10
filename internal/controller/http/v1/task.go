package v1

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strings"

	"github.com/ferralucho/go-task-management/internal/entity"
	"github.com/ferralucho/go-task-management/internal/usecase"
	"github.com/ferralucho/go-task-management/pkg/logger"
	"github.com/gin-gonic/gin"
)

type taskRoutes struct {
	t usecase.Card
	l logger.Interface
}

func newTaskRoutes(handler *gin.RouterGroup, t usecase.Card, l logger.Interface) {
	r := &taskRoutes{t, l}

	h := handler.Group("/management")
	{
		h.POST("/task", r.doCreateTask)
	}
}

type createTaskRequest struct {
	Type string `json:"type,required"`
}

type taskResponse struct {
	Name      string   `json:"name,required"`
	Desc      string   `json:"desc"`
	IdList    string   `json:"idList,required"`
	IdLabels  []string `json:"idLabels,required"`
	ShortLink string   `json:"shortLink"`
	ShortURL  string   `json:"shortUrl"`
	URL       string   `json:"url"`
}

// @Description Create task
// @ID          do-task
// @Tags  	    task
// @Accept      json
// @Produce     json
// @Param       request body doCreateTask true "Create task for management"
// @Success     200 {object} taskResponse
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /task [post]
func (r *taskRoutes) doCreateTask(c *gin.Context) {
	var request createTaskRequest
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		r.l.Error(err, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	var card entity.Card
	var err error

	switch entity.TaskType(strings.ToLower(request.Type)) {
	case entity.TypeTask:
		card, err = processTask(c, r, card, err)
	case entity.TypeIssue:
		card, err = processIssue(c, r, card, err)
	case entity.TypeBug:
		card, err = processBug(c, r, card, err)
	default:
		errorResponse(c, http.StatusBadRequest, "invalid request type")
		return
	}

	if err == nil {
		c.JSON(http.StatusOK, card)
	}
}

func processBug(c *gin.Context, r *taskRoutes, card entity.Card, err error) (entity.Card, error) {
	var bug entity.Bug
	if e := c.ShouldBindBodyWith(&bug, binding.JSON); e != nil {
		r.l.Error(e, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusBadRequest, "invalid bug request")
		return entity.Card{}, errors.New("invalid bug request")
	}

	card, err = r.t.CreateBug(c.Request.Context(), bug)
	if err != nil {
		r.l.Error(err, "http - v1 - processBug")
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return entity.Card{}, errors.New("invalid bug request")
	}
	return card, err
}

func processIssue(c *gin.Context, r *taskRoutes, card entity.Card, err error) (entity.Card, error) {
	var issue entity.Issue
	if e := c.ShouldBindBodyWith(&issue, binding.JSON); e != nil {
		r.l.Error(e, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusBadRequest, "invalid issue request")
		return entity.Card{}, errors.New("invalid issue request")
	}

	card, err = r.t.CreateIssue(c.Request.Context(), issue)

	if err != nil {
		r.l.Error(err, "http - v1 - processIssue")
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return entity.Card{}, errors.New("invalid bug request")
	}
	return card, err
}

func processTask(c *gin.Context, r *taskRoutes, card entity.Card, err error) (entity.Card, error) {
	var task entity.Task
	if e := c.ShouldBindBodyWith(&task, binding.JSON); e != nil {
		r.l.Error(e, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusBadRequest, "invalid task request")
		return entity.Card{}, errors.New("invalid task request")
	}

	card, err = r.t.CreateTask(c.Request.Context(), task)
	if err != nil {
		r.l.Error(err, "http - v1 - processTask")
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return entity.Card{}, errors.New("invalid bug request")
	}

	return card, err
}
