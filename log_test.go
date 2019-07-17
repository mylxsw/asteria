package asteria_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/mylxsw/asteria"
	"github.com/mylxsw/asteria/formatter"
	"github.com/mylxsw/asteria/level"
	"github.com/mylxsw/asteria/writer"
)

func TestModule(t *testing.T) {
	// DefaultLogLevel(Critical)

	// loc, _ := time.LoadLocation("Asia/Chongqing")
	// asteria.DefaultLocation(loc)
	// asteria.DefaultWithColor(false)
	// asteria.DefaultWithFileLine(true)

	asteria.AddGlobalFilter(func(filter asteria.Filter) asteria.Filter {
		return func(f formatter.Format) {
			// if f.Level == level.Debug {
			// 	return
			// }

			f.Context.UserContext["user_id"] = 123

			filter(f)
		}
	})

	asteria.GlobalContext(func(c formatter.LogContext) {
		c.SysContext["ref"] = "190101931"
	})
	asteria.Default().AddFilter(func(filter asteria.Filter) asteria.Filter {
		return func(f formatter.Format) {
			f.Context.SysContext["ref"] = "6f96cfdfe"
			filter(f)
		}
	})

	asteria.Default().LogLevel(level.Debug)
	asteria.Debug("细雨微风岸，危樯独夜舟")
	asteria.Error("月上柳梢头，人约黄昏后")
	asteria.WithContext(asteria.C{
		"user_id":  123,
		"username": "Tom",
	}).Warningf("君子坦荡荡，小人常戚戚")

	asteria.Module("order.test.scheduler").Noticef("order %s created", "1234592")
	asteria.Module("order.scheduler.module1.module2").Infof("order %s created", "1234592")
	asteria.Module("order.apple").Debugf("order %s created", "1234592")
	asteria.Module("order").Errorf("order %s created", "1234592")
	asteria.Module("order").Emergencyf("order %s created", "1234592")
	asteria.Module("order").Warningf("order %s created", "1234592")
	asteria.Module("order").Alertf("order %s created", "1234592")
	asteria.Module("order").Criticalf("order %s created\n", "1234592")

	asteria.Module("user").Formatter(formatter.NewJSONWithTimeFormatter()).Error("user create failed")

	asteria.WithContext(nil).Debug("error occur")
	asteria.Module("purchase").Formatter(formatter.NewJSONWithTimeFormatter()).WithContext(nil).Infof("用户 %s 已创建", "mylxsw")

	userLog := asteria.Module("user")
	userLog.WithContext(asteria.C{
		"id":   123,
		"name": "lixiaoyao",
	}).Debugf("Hello, %s", "world")

	taskLogger := asteria.Module("asteria.user.tasks").WithFileLine(true).GlobalContext(func(c formatter.LogContext) {
		asteria.GetDefaultConfig().GlobalContext(c)
		c.SysContext["enterprise_id"] = 15
	})
	taskLogger.Debug("Hello, world")

	enterpriseJobLogger := asteria.Module("asteria.user.enterprise.jobs").WithFileLine(true).Formatter(formatter.NewJSONFormatter())

	enterpriseJobLogger.AddFilter(func(filter asteria.Filter) asteria.Filter {
		return func(f formatter.Format) {
			// filter(f)
			f.Level = level.Emergency
			filter(f)
		}
	})

	enterpriseJobLogger.AddFilter(func(filter asteria.Filter) asteria.Filter {
		return func(f formatter.Format) {
			filter(f)
		}
	})

	enterpriseJobLogger.Debug("He remembered the count of Monte cristo")
	enterpriseJobLogger.Info("You are mistaken💯 I am not the Count of Monte Cristo")
	enterpriseJobLogger.Error("The noiseless door again turned on its hinges, and the Count of Monte Cristo reappeared")
	enterpriseJobLogger.WithContext(asteria.C{
		"user_id":  123,
		"username": "Tom",
	}).Warningf("君子坦荡荡，小人常戚戚")
}

func TestConcurrentWrite(t *testing.T) {
	var logger = asteria.Module("test.concurrent")

	var logfile = "./test.log"

	logger.Writer(writer.NewSingleFileWriter(logfile))

	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			logger.Debugf("inner - %d（%d）", i, rand.Intn(10))
		}(i)
	}

	for i := 0; i < 1000; i++ {
		logger.Debugf("outer - %d", i)
	}

	os.Remove(logfile)
}
