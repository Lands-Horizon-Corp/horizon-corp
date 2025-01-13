package controllers

import "github.com/gin-gonic/gin"

type FeedbackController struct{}

func NewFeedbackController() *FeedbackController {
	return &FeedbackController{}
}

// GET: /api/v1/feedback
// Retrieve feedback with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Allowed
//	Employee: Not allowed
//	Owner: Not allowed
//	Member: Not allowed
func (c *FeedbackController) Index(ctx *gin.Context) {}

// GET: /api/v1/feedback/:id
//
//	Admin: Allowed
//	Employee: Not allowed
//	Owner: Not allowed
//	Member: Not allowed
//
// Endpoint: GET /api/v1/feedback/:id
func (c *FeedbackController) Show(ctx *gin.Context) {}

// PUBLIC
// Endpoint: POST /api/v1/feedback
func (c *FeedbackController) Store(ctx *gin.Context) {}

//	Admin: Allowed
//	Employee: Not allowed
//	Owner: Not allowed
//	Member: Not allowed
//
// Endpoint: DELETE /api/v1/feedback/:id
func (c *FeedbackController) Destroy(ctx *gin.Context) {}
