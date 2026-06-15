package api

import (
	pb "bks/proto/users"

	"github.com/gin-gonic/gin"
)

var Client pb.UsersClient

func UserAdd(c *gin.Context) {
	var req pb.UsersAddReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := Client.UsersAdd(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func UserLogin(c *gin.Context) {
	var req pb.UsersLoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := Client.UsersLogin(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func UserInfo(c *gin.Context) {
	var req pb.UsersInfoReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := Client.UsersInfo(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func UserList(c *gin.Context) {
	var req pb.UsersListReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := Client.UsersList(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func UserUpdate(c *gin.Context) {
	var req pb.UsersUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := Client.UsersUpdate(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func UserDelete(c *gin.Context) {
	var req pb.UsersDeleteReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := Client.UsersDelete(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
