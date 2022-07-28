package main

import (
	"ginAndVueBBS/common"
	"ginAndVueBBS/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
)

func main() {
	r := gin.Default()
	InitConfig()
	common.InitDB()

	defer common.DB.Close()
	r = router.CollectRoute(r)
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImV4cCI6MTY1ODAzODUyMywiaWF0IjoxNjU3NDMzNzIzLCJpc3MiOiJ5dWFuc2hhbi5jb20iLCJzdWIiOiJ1c2VyIHRva2VuIn0.eMSk7FudW-zwjxFN_C3D_LGgcBbd593h-Lwixx7TS_A
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjMsImV4cCI6MTY1ODAzODUyMywiaWF0IjoxNjU3NDMzNzIzLCJpc3MiOiJ5dWFuc2hhbi5jb20iLCJzdWIiOiJ1c2VyIHRva2VuIn0.eMSk7FudW-zwjxFN_C3D_LGgcBbd593h-Lwixx7TS_A
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjAsImV4cCI6MTY1ODQxMTgyMSwiaWF0IjoxNjU3ODA3MDIxLCJpc3MiOiJ5dWFuc2hhbi5jb20iLCJzdWIiOiJ1c2VyIHRva2VuIn0.4nk5Cg--_nbICWfamodlUIgu3uBwPz3nUY5RfJ7hiDQ
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjAsImV4cCI6MTY1ODQxMjI1NiwiaWF0IjoxNjU3ODA3NDU2LCJpc3MiOiJ5dWFuc2hhbi5jb20iLCJzdWIiOiJ1c2VyIHRva2VuIn0.yZeCZ57kR6TBDguYHUk7RSfme6p76_T1SCe5kExBmMc
	panic(r.Run(":1016"))
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("initConfig fail err: " + err.Error())
	}
}
