// Package env provides a thin wrapper around accessing environment
// variables.
package env

import (
	"sync"
)

var muGlobal sync.RWMutex
var global Config

// SetPrefix sets the prefix used to format environment variable names
func SetPrefix(prefix string) *Config {
	muGlobal.Lock()
	defer muGlobal.Unlock()
	return global.SetPrefix(prefix)
}

// VarName formats the given name into an environment variable name
// using the rules defined by the global Config.
func VarName(name string) string {
	muGlobal.RLock()
	defer muGlobal.RUnlock()
	return global.VarName(name)
}

// Lookup is the same as `os.LookupEnv`, except it formats the
// environment variable named using `VarName()` before passing
// it to `os.LookupEnv`
func Lookup(name string) (string, bool) {
	muGlobal.RLock()
	defer muGlobal.RUnlock()
	return global.Lookup(name)
}

// Value returns the value of the environment variable named `name`,
// except it retrieves the value using `Lookup()`. Also, if the
// value of the environment variable is empty, it returns `defaultValue`
func Value(name, defaultValue string) string {
	muGlobal.RLock()
	defer muGlobal.RUnlock()
	return global.Value(name, defaultValue)
}

// SetValue sets the value of the environment variable named `name`
// to `value`, except it formats the environment variable name using
// `VarName()` before passing it to `os.Setenv`
func SetValue(name, value string) error {
	muGlobal.RLock()
	defer muGlobal.RUnlock()
	return global.SetValue(name, value)
}
