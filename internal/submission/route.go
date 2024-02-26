package submission

import "github.com/gin-gonic/gin"

func SubmissionRoute(engine *gin.Engine) {
	router := engine.Group("/submissions")
	router.POST("/")
}
