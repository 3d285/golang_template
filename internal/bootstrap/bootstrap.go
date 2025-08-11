package bootstrap

import (
	"fmt"
	"net/http"
	"golang_template/internal/config"
)

func Run() error {
	if err := config.Load("configs/config.yaml"); err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	return http.ListenAndServe(":8080", nil)
}
