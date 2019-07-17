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

	asteria.GlobalContext(func(c formatter.LogContext) {
		c.SysContext["ref"] = "190101931"
	})

	asteria.GetDefaultModule().LogLevel(level.Debug)
	asteria.Debug("xxxx")

	asteria.Module("order.test.scheduler").Noticef("order %s created", "1234592")
	asteria.Module("order.scheduler.module1.module2").Infof("order %s created", "1234592")
	asteria.Module("order.xajckakejcjakjk").Debugf("order %s created", "1234592")
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
		asteria.Default().GlobalContext(c)
		c.SysContext["enterprise_id"] = 15
	})
	taskLogger.Debug("Hello, world")
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
