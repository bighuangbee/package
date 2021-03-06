package request

import (
	"github.com/gin-gonic/gin"
	"gopackage/http/respone"
)

func ShouldBind(c *gin.Context, obj interface{}){
	if err := c.ShouldBind(obj); err != nil {
		respone.Error(c, respone.INVALID_PARAMS, err.Error())
	}
}

func BindJSON(c *gin.Context, obj interface{}) error{
	if err := c.ShouldBindJSON(obj); err != nil {
		respone.Error(c, respone.INVALID_PARAMS, err.Error())
		return err
	}
	return nil
}