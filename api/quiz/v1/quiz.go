package v1

import (
	"math/rand"
	"strconv"

	"github.com/KatsuyaKawabe/go-quiz-api/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Create(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "quiz create"})
	return
}

// GET /api/v1/quiz?difficulty=1,2,3
func Index(ctx *gin.Context) {
	qS := ctx.Query("difficulty")
	parseQS, _ := strconv.ParseUint(qS, 10, 32)
	difficulty := uint(parseQS)

	var qs []models.Quiz
	sess, _ := mgo.Dial(viper.GetString("mongo.host"))
	defer sess.Close()
	db := sess.DB(viper.GetString("mongo.db"))
	col := db.C("quiz")

	//  難易度指定でいくつかのクイズを取得
	colCount, err := col.Find(bson.M{"difficulty": difficulty}).Count()
	if err != nil {
		ctx.JSON(400, "Invalid request")
		return
	}

	if colCount == 0 {
		ctx.JSON(404, "Quiz Not found")
		return
	}

	query := col.Find(bson.M{"difficulty": difficulty})
	query.All(&qs)
	quiz := qs[rand.Intn(colCount)]
	ctx.JSON(200, quiz)
	return
}
