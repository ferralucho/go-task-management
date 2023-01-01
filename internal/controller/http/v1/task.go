package v1

import (
	"net/http"
	"strings"

	"github.com/ferralucho/go-task-management/internal/entity"
	"github.com/ferralucho/go-task-management/internal/usecase"
	"github.com/ferralucho/go-task-management/pkg/logger"
	"github.com/gin-gonic/gin"
)

type taskRoutes struct {
	t usecase.CardUseCase
	l logger.Interface
}

func newTaskRoutes(handler *gin.RouterGroup, t usecase.CardUseCase, l logger.Interface) {
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
	if err := c.ShouldBindJSON(&request); err != nil {
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

	if err != nil {
		r.l.Error(err, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusInternalServerError, "tasks service problem")

		return
	}

	c.JSON(http.StatusOK, card)
}

func processBug(c *gin.Context, r *taskRoutes, card entity.Card, err error) (entity.Card, error) {
	var bug entity.Bug
	if e := c.ShouldBindJSON(&bug); e != nil {
		r.l.Error(e, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusBadRequest, "invalid request bug")

		return entity.Card{}, nil
	}

	card, err = r.t.CreateBug(c.Request.Context(), bug)
	return card, err
}

func processIssue(c *gin.Context, r *taskRoutes, card entity.Card, err error) (entity.Card, error) {
	var issue entity.Issue
	if e := c.ShouldBindJSON(&issue); e != nil {
		r.l.Error(e, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusBadRequest, "invalid request issue")

		return entity.Card{}, nil
	}

	card, err = r.t.CreateIssue(c.Request.Context(), issue)
	return card, err
}

func processTask(c *gin.Context, r *taskRoutes, card entity.Card, err error) (entity.Card, error) {
	var task entity.Task
	if e := c.ShouldBindJSON(&task); e != nil {
		r.l.Error(e, "http - v1 - doCreateTask")
		errorResponse(c, http.StatusBadRequest, "invalid request task")

		return entity.Card{}, nil
	}

	card, err = r.t.CreateTask(c.Request.Context(), task)
	return card, err
}
