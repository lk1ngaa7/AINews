package helpers

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitResource()
	m.Run()
	os.Exit(0)
}
