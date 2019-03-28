package callerstack

import (
	"strings"
	"testing"

	"github.com/xgfone/logger"
)

func TestCaller(t *testing.T) {
	if v, err := Caller()(logger.Record{Lvl: logger.LvlDebug}); err != nil {
		t.Error(err)
	} else if v == nil {
		t.Fail()
	} else if s, ok := v.(string); !ok || !strings.HasSuffix(s, ":11") {
		t.Error(v)
	}
}

func TestCallerStack(t *testing.T) {
	if v, err := CallerStack()(logger.Record{Lvl: logger.LvlDebug}); err != nil {
		t.Error(err)
	} else if v == nil {
		t.Fail()
	} else if s, ok := v.(string); !ok || !strings.HasSuffix(s, ":21]") {
		t.Error(v)
	}
}
