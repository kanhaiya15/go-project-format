package gopfserver

import (
	"github.com/kanhaiya15/gopf/utils"
)

// Init main function for server start
func Init() {
	r := NewRouter()
	port, err := utils.GetConfValue("APP_PORT")
	if err != nil {
		panic(err.Error())
	}
	r.Run(":" + port)
}
