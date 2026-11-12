package mylog

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestInfoOutput(t *testing.T) {
	out := captureOutput(func() {
		Info("Hello %s", "Jenderal")
	})

	// Cek prefix dan pesan secara terpisah (karena ada ANSI di antaranya)
	if !strings.Contains(out, "[INF]") {
		t.Errorf("Prefix [INF] tidak ditemukan: %q", out)
	}
	if !strings.Contains(out, "Hello Jenderal") {
		t.Errorf("Pesan tidak ditemukan: %q", out)
	}
	if !strings.Contains(out, colorCyan) {
		t.Errorf("Warna Cyan tidak ditemukan: %q", out)
	}
}

func TestLevelFilter(t *testing.T) {
	SetLevel(ERROR)
	defer SetLevel(INFO)

	out := captureOutput(func() {
		Info("Pesan ini harus disembunyikan")
		Error("Pesan ini harus muncul")
	})

	if strings.Contains(out, "disembunyikan") {
		t.Errorf("Level filter gagal, pesan INFO ikut tercetak")
	}
	if !strings.Contains(out, "Pesan ini harus muncul") {
		t.Errorf("Level filter gagal, pesan ERROR tidak tercetak")
	}
}

func TestTimestampOff(t *testing.T) {
	SetTimestamp(false)
	defer SetTimestamp(true)

	out := captureOutput(func() {
		Warn("Tanpa Waktu")
	})

	if !strings.HasPrefix(out, colorYellow+"[WRN]") {
		t.Errorf("Timestamp gagal dimatikan. Output: %s", out)
	}
}

func TestBoldPrefix(t *testing.T) {
	SetBold(true)
	defer SetBold(false)

	out := captureOutput(func() {
		Info("Test bold")
	})

	if !strings.Contains(out, colorBold) {
		t.Errorf("Bold code tidak muncul di output: %q", out)
	}
}
