package config

import (
	"os"
)

func Load(path string) error {
	// demo: pretend to read; ensure file exists if provided
	if path == "" { return nil }
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// it's fine for the demo; in real code, parse YAML/JSON here
		return nil
	}
	return nil
}
