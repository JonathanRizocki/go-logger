package main

import (
	"os"
	"time"

	"learngo-pockets/logger/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout, 1000))

	lgr.LogF(pocketlog.LevelInfo, "A little copying is better than a little dependency.")
	lgr.LogF(pocketlog.LevelError, "Errors are values. Documentation is for %s.", "users")
	lgr.LogF(pocketlog.LevelDebug, "Make the zero (%d) value useful.", 0)

	lgr.LogF(pocketlog.LevelInfo, "Hallo, %d %v", 2024, time.Now())
}
