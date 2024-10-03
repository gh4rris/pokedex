package main

import "os"

func commandExit(cfg *config, area string) error {
	os.Exit(0)
	return nil
}
