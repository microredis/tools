package echoutil

import "net/http"

func StatusWithStatusText(status int) (int, string) {
	return status, http.StatusText(status)
}
