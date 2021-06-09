package service

import (
	"context"
	"log"
	"logger/app/logdriver"
	"logger/app/model"
	"logger/lib/hook"
	"logger/lib/snsq"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	LogTable   = "log"
	LogTime    = "log_time"
	LogSys     = "log_sys"
	LogMsg     = "log_msg"
	LogLevel   = "log_level"
	LogContent = "content"
	LogHost    = "host"
	LogTopic   = "topic"
	LogVersion = "version"
	// mongo field
)

// MongoLogHook，继承基础钩子
type MongoLogHook struct {
	*hook.BashHook
}

// NewMongoHook
func NewMongoHook(level logrus.Level, h hook.HookFireFunc) *MongoLogHook {
	return &MongoLogHook{hook.NewBashHook(level, h)}
}

// mongoSyncFireFunc 实际逻辑操作,入库定义,func(host string, entry *logrus.Entry) error
func (s *Service) mongoSyncFireFunc(entry *logrus.Entry) error {
	m := fieldsToBson(entry)
	if _, err := s.MongoDB.Collection(LogTable).InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}
func (s *Service) MongoHooKInit(driver, database string) logrus.Hook {
	mongo, err := logdriver.OpenMongo(driver, database)
	if err != nil {
		panic(err)
	}
	s.MongoDB = mongo
	lh := NewMongoHook(logrus.DebugLevel, s.mongoSyncFireFunc)
	return lh
}

func fieldsToBson(entry *logrus.Entry) bson.M {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	m := bson.M{LogSys: entry.Message, LogLevel: entry.Level}
	for _, val := range entry.Data {
		v, ok := val.(snsq.NsqMesContent)
		if !ok {
			panic("fieldsToBson")
		}
		m[LogContent] = v.Content
		m[LogHost] = v.Host
		m[LogTopic] = v.Topic
		m[LogVersion] = v.Version
		m[LogTime] = v.LogTime
	}

	return m
}

func (s *Service) Find(ArgCondition *model.ArgLogCondition) ([]bson.M, error) {
	res := []bson.M{}
	condition := contionToBson(ArgCondition)
	opts := getFindOptions(ArgCondition)
	cur, err := s.MongoDB.Collection(LogTable).Find(context.Background(), condition, opts...)
	if err != nil {
		return res, err
	}

	for cur.Next(context.Background()) {
		var p bson.M
		if err := cur.Decode(&p); err != nil {
			return res, err
		}
		res = append(res, p)
	}
	return res, nil
}

func getFindOptions(ArgCondition *model.ArgLogCondition) []*options.FindOptions {
	if ArgCondition.Page == 0 {
		ArgCondition.Page = 1
	}
	if ArgCondition.PageSize == 0 {
		ArgCondition.Page = 20
	}
	opts := []*options.FindOptions{}
	limit := options.Find().SetLimit(ArgCondition.PageSize)
	opts = append(opts, limit)
	skip := options.Find().SetSkip((ArgCondition.Page - 1) * ArgCondition.PageSize)
	opts = append(opts, skip)

	if ArgCondition.Order.OrderField != "" && ArgCondition.Order.OrderVal != 0 {
		order := options.Find().SetSort(bson.M{ArgCondition.Order.OrderField: ArgCondition.Order.OrderVal})
		opts = append(opts, order)
	}
	return opts
}

func contionToBson(arg *model.ArgLogCondition) bson.M {
	res := bson.M{}
	toCheckValues("log_sys", "", arg.LogSys, res)
	toCheckValues("log_level", "$in", arg.LogLevel, res)
	toCheckValues("topic", "$regex", arg.Topic, res)
	toCheckValues("content", "$regex", arg.Content, res)
	if !arg.LogStartTime.IsZero() {
		res["log_time"] = bson.M{"$gte": arg.LogStartTime}
	}
	if !arg.LogEndTime.IsZero() {
		res["log_time"] = bson.M{"$lte": arg.LogEndTime}
	}
	return res
}

func toCheckValues(field, identifier, content string, condition bson.M) {
	switch {
	case content == "":
		return
	case identifier == "":
		condition[field] = content
	case identifier == "$in":
		levels := strings.Split(content, ",")
		list := []int{}
		for _, v := range levels {
			l, err := strconv.Atoi(v)
			if err != nil {
				return
			}
			list = append(list, l)
		}
		condition[field] = bson.M{"$in": list}
	case identifier == "$regex":
		condition[field] = bson.M{"$regex": primitive.Regex{Pattern: "[" + content + "]+", Options: "im"}}
	}
}
func getOrder(order []model.ArgOrder) *options.FindOptions {
	b := bson.M{}
	for _, v := range order {
		b[v.OrderField] = v.OrderVal
	}
	return options.Find().SetSort(b)

}
