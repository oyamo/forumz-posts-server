package handlers

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"posts-server/internal/domain/posts"
	"posts-server/internal/interfaces/web/dto"
)

type PostsHandler struct {
	postsUC   *posts.UseCase
	logger    *zap.SugaredLogger
	validator *validator.Validate
}

func (h *PostsHandler) CreatePost(c *gin.Context) {
	var responseDto dto.ResponseDto
	requestIdCtx, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{})
		h.logger.Errorw("cannot find id from context")
		return
	}

	requestId, ok := requestIdCtx.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		h.logger.Errorw("requestId is not uuid type")
		return
	}

	responseDto.RequestId = requestId
	var req posts.CreatePostDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		responseDto.Description = "Missing/invalid request parameters"
		c.JSON(http.StatusBadRequest, responseDto)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		h.logger.Errorw("error validating request", "error", err)
		c.JSON(http.StatusBadRequest, responseDto)
		return
	}

	initiator, exists := c.Get("initiator")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{})
		h.logger.Errorw("cannot find id from context")
		return
	}

	initiatorId, ok := initiator.(uuid.UUID)
	if !ok {
		responseDto.Description = "Something went wrong"
		c.JSON(http.StatusInternalServerError, responseDto)
		h.logger.Errorw("initiator is not uuid type")
		return
	}

	req.PersonId = initiatorId

	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "id", responseDto.RequestId)

	post, err := h.postsUC.Create(ctx, req)
	if err != nil {
		responseDto.Description = "Something went wrong"
		c.JSON(http.StatusInternalServerError, responseDto)
		h.logger.Errorw("error creating post", "error", err)
		return
	}

	responseDto.Description = "Post successfully created"
	responseDto.Data = post
	c.JSON(http.StatusCreated, responseDto)

}

func (h *PostsHandler) DeletePost(c *gin.Context) {
	var responseDto dto.ResponseDto
	requestIdCtx, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{})
		h.logger.Errorw("cannot find id from context")
		return
	}

	requestId, ok := requestIdCtx.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		h.logger.Errorw("requestId is not uuid type")
		return
	}

	responseDto.RequestId = requestId
	postId, err := uuid.Parse(c.Param("postId"))
	if err != nil {
		responseDto.Description = "Post Id is invalid"
		c.JSON(http.StatusBadRequest, responseDto)
		return
	}

	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "id", responseDto.RequestId)

	post, err := h.postsUC.Find(ctx, postId)
	if err != nil {
		msg := "Something went wrong"
		if errors.Is(err, sql.ErrNoRows) {
			msg = "Post not found"
		} else {
			h.logger.Errorw("error finding post", "error", err)
		}
		responseDto.Description = msg
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}

	responseDto.Description = "Success"
	responseDto.Data = post
	c.JSON(http.StatusOK, responseDto)
}

func (h *PostsHandler) ViewPost(c *gin.Context) {
	var responseDto dto.ResponseDto
	requestIdCtx, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{})
		h.logger.Errorw("cannot find id from context")
	}

	requestId, ok := requestIdCtx.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		h.logger.Errorw("requestId is not uuid type")
		return
	}

	responseDto.RequestId = requestId
	postId, err := uuid.Parse(c.Param("postId"))
	if err != nil {
		responseDto.Description = "Post Id is invalid"
		c.JSON(http.StatusBadRequest, responseDto)
		return
	}

	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "id", responseDto.RequestId)

	post, err := h.postsUC.Find(ctx, postId)
	if err != nil {
		responseDto.Description = "Something went wrong"
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, sql.ErrNoRows):
			responseDto.Description = "Post not found/deleted"
			status = http.StatusNotFound
		}
		c.JSON(status, responseDto)
		return
	}

	responseDto.Description = "Post successfully found"
	responseDto.Data = post
	c.JSON(http.StatusOK, responseDto)
}

func NewPostsHandler(postsUC *posts.UseCase, logger *zap.SugaredLogger) *PostsHandler {
	return &PostsHandler{
		postsUC:   postsUC,
		logger:    logger,
		validator: validator.New(),
	}
}
