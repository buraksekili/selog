package selog_test

import (
	"bytes"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/buraksekili/selog"
)

func TestSelog_Debug(t *testing.T) {

	t.Run("DEBUG without formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Debug("testing debug without formatting")

		str := buf.String()
		pi := strings.Index(str, "[DEBUG]")

		assert.Equal(t, "[DEBUG]", str[pi:pi+7])
		assert.Equal(t, "[DEBUG] testing debug without formatting", str[pi:])
	})

	t.Run("DEBUG with formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		state := "debug"
		s.Debug("testing %s with formatting %s!\n", state, "Burak")

		str := buf.String()
		pi := strings.Index(str, "[DEBUG]")
		assert.Equal(t, "[DEBUG]", str[pi:pi+7])
		assert.Equal(t, "[DEBUG] testing debug with formatting Burak!\n", str[pi:])
	})

	t.Run("DEBUG with empty message", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Debug("")

		str := buf.String()
		pi := strings.Index(str, "[DEBUG]")

		assert.Equal(t, "[DEBUG]", str[pi:pi+7])
		assert.Equal(t, "[DEBUG] ", str[pi:])
	})

}

func TestSelog_Info(t *testing.T) {
	t.Run("INFO without formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Info("testing INFO without formatting")

		str := buf.String()
		l := len("[INFO]")
		pi := strings.Index(str, "[INFO]")

		assert.Equal(t, "[INFO]", str[pi:pi+l])
		assert.Equal(t, "[INFO] testing INFO without formatting", str[pi:])
	})

	t.Run("INFO with formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		state := "INFO"
		s.Info("testing %s with formatting %s!\n", state, "Burak")

		str := buf.String()
		l := len("[INFO]")
		pi := strings.Index(str, "[INFO]")

		assert.Equal(t, "[INFO]", str[pi:pi+l])
		assert.Equal(t, "[INFO] testing INFO with formatting Burak!\n", str[pi:])
	})

	t.Run("INFO with empty message", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Info("")

		str := buf.String()
		l := len("[INFO]")
		pi := strings.Index(str, "[INFO]")

		assert.Equal(t, "[INFO]", str[pi:pi+l])
		assert.Equal(t, "[INFO] ", str[pi:])
	})
}

func TestSelog_Error(t *testing.T) {
	t.Run("ERROR without formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Error("testing ERROR without formatting")

		str := buf.String()
		l := len("[ERROR]")
		strings.Trim(str, "\\x1b[0m")
		pi := strings.Index(str, "[ERROR]")

		assert.Equal(t, "[ERROR]", str[pi:pi+l])
		assert.Equal(t, "[ERROR] testing ERROR without formatting", str[pi:])
	})

	t.Run("ERROR with formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		state := "ERROR"
		s.Error("testing %s with formatting %s!\n", state, "Burak")

		str := buf.String()
		l := len("[ERROR]")
		strings.Trim(str, "\\x1b[0m")
		pi := strings.Index(str, "[ERROR]")

		assert.Equal(t, "[ERROR]", str[pi:pi+l])
		assert.Equal(t, "[ERROR] testing ERROR with formatting Burak!\n", str[pi:])
	})

	t.Run("ERROR with empty message", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Error("")

		str := buf.String()
		l := len("[ERROR]")
		pi := strings.Index(str, "[ERROR]")

		assert.Equal(t, "[ERROR]", str[pi:pi+l])
		assert.Equal(t, "[ERROR] ", str[pi:])
	})
}

//

func TestSelog_Success(t *testing.T) {
	t.Run("SUCCESS without formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Success("testing SUCCESS without formatting")
		str := buf.String()
		l := len("[SUCCESS]")
		pi := strings.Index(str, "[SUCCESS]")

		assert.Equal(t, "[SUCCESS]", str[pi:pi+l])
		assert.Equal(t, "[SUCCESS] testing SUCCESS without formatting", str[pi:])
	})

	t.Run("SUCCESS with formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		state := "SUCCESS"
		s.Success("testing %s with formatting %s!\n", state, "Burak")

		str := buf.String()
		l := len("[SUCCESS]")
		pi := strings.Index(str, "[SUCCESS]")

		assert.Equal(t, "[SUCCESS]", str[pi:pi+l])
		assert.Equal(t, "[SUCCESS] testing SUCCESS with formatting Burak!\n", str[pi:])
	})

	t.Run("SUCCESS with empty message", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Success("")

		str := buf.String()
		l := len("[SUCCESS]")
		pi := strings.Index(str, "[SUCCESS]")

		assert.Equal(t, "[SUCCESS]", str[pi:pi+l])
		assert.Equal(t, "[SUCCESS] ", str[pi:])
	})
}

func TestSelog_Warn(t *testing.T) {
	t.Run("WARN without formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Warn("testing WARN without formatting")

		str := buf.String()
		l := len("[WARN]")
		pi := strings.Index(str, "[WARN]")

		assert.Equal(t, "[WARN]", str[pi:pi+l])
		assert.Equal(t, "[WARN] testing WARN without formatting", str[pi:])
	})

	t.Run("WARN with formatting", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		state := "WARN"
		s.Warn("testing %s with formatting %s!\n", state, "Burak")

		str := buf.String()
		l := len("[WARN]")
		pi := strings.Index(str, "[WARN]")

		assert.Equal(t, "[WARN]", str[pi:pi+l])
		assert.Equal(t, "[WARN] testing WARN with formatting Burak!\n", str[pi:])
	})

	t.Run("WARN with empty message", func(t *testing.T) {
		var buf bytes.Buffer
		s := getNewSelog(&buf)

		s.Warn("")

		str := buf.String()
		l := len("[WARN]")
		pi := strings.Index(str, "[WARN]")

		assert.Equal(t, "[WARN]", str[pi:pi+l])
		assert.Equal(t, "[WARN] ", str[pi:])
	})
}

func getNewSelog(w io.Writer) *selog.Selog {
	l := log.New(w, "TEST", log.LstdFlags)
	return selog.NewLogger(l)
}

// func TestSelog_Fatal(t *testing.T) {
// 	t.Run("FATAL without formatting", func(t *testing.T) {
// 		var buf bytes.Buffer
// 		s := getNewSelog(&buf)
//
// 		s.Fatal("testing FATAL without formatting")
//
// 		str := buf.String()
// 		l := len("[FATAL]")
// 		pi := strings.Index(str, "[FATAL]")
//
// 		assert.Equal(t, "[FATAL]", str[pi:pi+l])
// 		assert.Equal(t, "[FATAL] testing FATAL without formatting", str[pi:])
// 	})
//
// 	t.Run("FATAL with formatting", func(t *testing.T) {
// 		var buf bytes.Buffer
// 		s := getNewSelog(&buf)
//
// 		state := "FATAL"
// 		s.Fatal("testing %s with formatting %s!\n", state, "Burak")
//
// 		str := buf.String()
// 		l := len("[FATAL]")
// 		pi := strings.Index(str, "[FATAL]")
//
// 		assert.Equal(t, "[FATAL]", str[pi:pi+l])
// 		assert.Equal(t, "[FATAL] testing FATAL with formatting Burak!\n", str[pi:])
// 	})
//
// 	t.Run("FATAL with empty message", func(t *testing.T) {
// 		var buf bytes.Buffer
// 		s := getNewSelog(&buf)
//
// 		s.Fatal("")
//
// 		str := buf.String()
// 		l := len("[FATAL]")
// 		pi := strings.Index(str, "[FATAL]")
//
// 		assert.Equal(t, "[FATAL]", str[pi:pi+l])
// 		assert.Equal(t, "[FATAL] ", str[pi:])
// 	})
// }
