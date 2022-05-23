package sql

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongSqlController struct {
	beego.Controller
}

type LogRecord struct {
	UserName string     `bson:"username"` //任务名
	PassWord string     `bson:"password"` //shell命令
}

func (c *MongSqlController) Get() {
	name := c.GetString("username")
	f := c.GetString("flag")
	uri := "mongodb://admin:123456@172.18.0.4:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	filter := bson.M{"username":name}

	//2.选择数据库 my_db里的某个表
	collection := client.Database("beego").Collection("users")
	var cursor *mongo.Cursor

	if f == "true"{
		singleResult := collection.FindOne(context.TODO(),filter)

		if singleResult.Err() != nil {
			log.Println("Find error: ", singleResult.Err())
		}

		UserData := LogRecord{}

		decodeError := singleResult.Decode(&UserData)

		if decodeError != nil {
			log.Println("Decode error: ", decodeError)
		}

		c.Data["userdata"] = UserData
	}else{
		if cursor, err = collection.Find(context.TODO(), filter, options.Find().SetSkip(0), options.Find().SetLimit(2)); err != nil {
			fmt.Println(err)
			return
		}
		//延迟关闭游标
		defer func() {
			if err = cursor.Close(context.TODO()); err != nil {
				panic(err)
			}
		}()

		var results []LogRecord
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}
		c.Data["results"] = results
	}

	c.TplName = "vuln/sql/Mongdb.tpl"
}