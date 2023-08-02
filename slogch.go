package chslog

import (
	"atomicgo.dev/isprod"
	"log/slog"
)

// Conditions holds the environment conditions that are checked.
// By default, it checks for the most common environments with https://atomicgo.dev/isprod.
// See https://pkg.go.dev/atomicgo.dev/isprod#Conditions for more info, on how to configure them (if needed).
var Conditions = isprod.DefaultConditions

// Choose automatically chooses between prodHandler and devHandler.
func Choose(prodHandler, devHandler slog.Handler) slog.Handler {
	if Conditions.Check() {
		return prodHandler
	} else {
		return devHandler
	}
}
