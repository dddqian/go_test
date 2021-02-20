package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"
)

func SysInfo(c *gin.Context) {
	sysinfo := map[string]interface{}{
		"系统类型：":  runtime.GOOS,
		"系统架构：":  runtime.GOARCH,
		"CPU核数：": runtime.NumCPU(),
	}

	hostname, err := os.Hostname()
	if err != nil {
		sysinfo["电脑名称："] = "未知"
	} else {
		sysinfo["电脑名称："] = hostname
	}

	c.JSON(http.StatusOK, sysinfo)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
