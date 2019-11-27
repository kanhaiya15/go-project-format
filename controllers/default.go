package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kanhaiya15/gopf/conf/dbs/gopfmysql"
	"github.com/kanhaiya15/gopf/logging/gopflogrus"
)

var (
	logger = gopflogrus.NewLogger()
)

// Stats DB Status
type Stats struct {
	DBstats interface{} `json:"dbstats"`
}

// Home Base
// @Router / [get]
func Home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hi"})
}

// Ping pong
// @Router /v1/ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// Health health
// @Router /v1/health [get]
func Health(c *gin.Context) {
	result := map[string]interface{}{}
	err := gopfmysql.DBstatus()
	if err != nil {
		result["master-db-status"] = "inactive"
		logger.Errorf("Error! BaseHealth database.DBstatus() error response %+v\n", err.Error())
	} else {
		result["master-db-status"] = "active"
	}
	c.JSON(200, result)
}

// DBStats stats
// @Router /v1/stats [get]
func DBStats(c *gin.Context) {
	stats := Stats{
		DBstats: gopfmysql.DBstats(),
	}
	resultjson, err := json.Marshal(stats)
	if err != nil {
		logger.Errorf("Error! GetDBStats json.Marshal(stats) error response %+v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"resultstats": string(resultjson),
		})
	}
}
