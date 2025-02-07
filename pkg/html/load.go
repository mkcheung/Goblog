package html

import "github.com/gin-gonic/gin"

func LoanHTML(router *gin.Engine) {
	router.LoadHTMLGlob("internal/**/**/**/*tmpl")
}
