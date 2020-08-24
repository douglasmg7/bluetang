package bluetang

import (
	"os"
	"testing"
	// "github.com/douglasmg7/bluetang"
)

func TestMain(m *testing.M) {

	setupTest()
	code := m.Run()
	shutdownTest()

	os.Exit(code)
}

func setupTest() {
}

func shutdownTest() {
}

// Create satment insert product.
func TestUint64Number(t *testing.T) {
	val, msg := Uint64Number("123")
	if val == 0 || msg != "" {
		t.Errorf("Uint64Number(\"123\") = %d, %s; want 123 and \"\"", val, msg)
	}
	val, msg = Uint64Number("-123")
	if val != 0 || msg != "Número inválido" {
		t.Errorf("Uint64Number(\"-123\") = %d, %s; want 0 and \"Número inválido\"", val, msg)
	}
}

// Create satment insert product.
func TestUint32Number(t *testing.T) {
	val, msg := Uint32Number("123")
	if val == 0 || msg != "" {
		t.Errorf("Uint32Number(\"123\") = %d, %s; want 123 and \"\"", val, msg)
	}
	val, msg = Uint32Number("-123")
	if val != 0 || msg != "Número inválido" {
		t.Errorf("Uint32Number(\"-123\") = %d, %s; want 0 and \"Número inválido\"", val, msg)
	}
}
