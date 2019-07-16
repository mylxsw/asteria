package asteria_test

import (
	"log/syslog"
	"testing"

	"github.com/mylxsw/asteria"
)

func TestSyslogWriter(t *testing.T) {
	logger := asteria.Module("syslog")
	logger.Writer(asteria.NewSyslogWriter("", "", syslog.LOG_INFO | syslog.LOG_LOCAL7, "asteria"))

	logger.Debug("As usual, a pilot put off immediately, and rounding the Chateau d'If.")
	logger.Critical("got on board the vessel between Cape Morgion and Rion island")
	logger.Info("Immediately, and according to custom, the ramparts of Fort Saint-Jean were covered with spectators")
	logger.Notice("Immediately, and according to custom, the ramparts of Fort Saint-Jean were covered with spectators")
	logger.Warning("Immediately, and according to custom, the ramparts of Fort Saint-Jean were covered with spectators")
	logger.Error("it is always an event at Marseilles for a ship to come into port")
	logger.Emergency("especially when this ship, like the Pharaon, has been built, rigged, and laden at the old Phocée docks, and belongs to an owner of the city")
}
