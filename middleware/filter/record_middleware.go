package filter

import (
	"net/http"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/middleware/provider"
	"github.com/VinncentWong/Delos-AquaFarm/util"
	"github.com/gin-gonic/gin"
)

type RecordMiddleware struct {
	provider provider.IRecordProvider
}

func NewRecordMiddleware(provider provider.IRecordProvider) *RecordMiddleware {
	return &RecordMiddleware{
		provider: provider,
	}
}

func (r *RecordMiddleware) SaveTrackedData() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipAddress := util.GetIpAddress(c)
		record := domain.RecordApi{
			Endpoint:   c.Request.URL.Path,
			IpAddress:  ipAddress,
			MethodName: c.Request.Method,
		}
		err := r.provider.SaveRecord(&record)
		if err != nil {
			c.Abort()
			util.SendResponse(c, http.StatusInternalServerError, err.Error(), false, nil)
			return
		}
		c.Next()
	}
}
