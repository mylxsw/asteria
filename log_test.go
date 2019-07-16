package asteria_test

import (
	"testing"
	"time"

	"asteria"
)

func TestModule(t *testing.T) {
	// DefaultLogLevel(LevelCritical)

	loc, _ := time.LoadLocation("Asia/Chongqing")
	asteria.DefaultLocation(loc)
	// asteria.DefaultWithColor(false)
	// asteria.DefaultWithFileLine(true)

	asteria.GlobalContext(func(c asteria.LogContext) {
		c.SysContext["ref"] = "190101931"
	})

	asteria.GetDefaultModule().LogLevel(asteria.LevelDebug)
	asteria.Debug("xxxx")

	asteria.Module("order.test.scheduler").Noticef("order %s created", "1234592")
	asteria.Module("order.scheduler.module1.module2").Infof("order %s created", "1234592")
	asteria.Module("order.xajckakejcjakjk").Debugf("order %s created", "1234592")
	asteria.Module("order").Errorf("order %s created", "1234592")
	asteria.Module("order").Emergencyf("order %s created", "1234592")
	asteria.Module("order").Warningf("order %s created", "1234592")
	asteria.Module("order").Alertf("order %s created", "1234592")
	asteria.Module("order").Criticalf("order %s created\n", "1234592")

	asteria.Module("user").Formatter(asteria.NewJSONFormatter()).Error("user create failed")

	asteria.WithContext(nil).Debug("error occur")
	asteria.Module("purchase").Formatter(asteria.NewJSONFormatter()).WithContext(map[string]interface{}{}).Infof("用户 %s 已创建", "mylxsw")

	userLog := asteria.Module("user")
	userLog.WithContext(asteria.C{
		"id":   123,
		"name": "lixiaoyao",
	}).Debugf("Hello, %s", "world")

	taskLogger := asteria.Module("asteria.user.tasks").WithFileLine(true).GlobalContext(func(c asteria.LogContext) {
		asteria.Default().GlobalContext(c)
		c.SysContext["enterprise_id"] = 15
	})
	taskLogger.Debug("Hello, world")
}
