package conf

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kanhaiya15/gopf/utils"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	runIn, err := utils.GetConfValue("APP_ENV")
	if err != nil {
		panic(err.Error())
	}
	runFile := wd + "/conf/environments/" + runIn + "/environment." + runIn + ".env"
	godotenv.Load(runFile)
}
