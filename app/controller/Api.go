package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tattoo_code/app/models"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// CreateInfo 创建信息
func CreateInfo(c *gin.Context) {
	// 打印请求的body
	fmt.Println(c.Request.Body)
	var info models.Info
	err := c.ShouldBindJSON(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = info.CreateInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, info)
}

// GetInfoById 根据id获取信息
func GetInfoById(c *gin.Context) {
	// 根据参数获取id
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var info models.Info
	info, err = info.GetInfoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"data":    info,
		"message": "获取成功",
	})
}

// UpdateInfo 修改信息
func UpdateInfo(c *gin.Context) {
	// 根据参数获取id
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var info models.Info
	info, err = info.GetInfoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.ShouldBindJSON(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = info.UpdateInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, info)
}

// DeleteInfo 删除信息
func DeleteInfo(c *gin.Context) {
	// 根据参数获取id
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var info models.Info
	info, err = info.GetInfoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = info.DeleteInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "删除成功",
	})
}
