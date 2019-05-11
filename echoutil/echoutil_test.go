package echoutil

import (
	"net/http"
	"testing"
)

func TestStatusWithStatusText(t *testing.T) {
	status, statusText := StatusWithStatusText(http.StatusOK)
	if status != http.StatusOK {
		t.Fatal()
	}
	if statusText != "OK" {
		t.Fatal()
	}
}
