package env

import (
	"os"
)

type Config struct {
	prefix    string
	separator string
}

// Separator returns the separator used to join the components
// of an environment variable. By defualt the separator is "_".
func (c *Config) Separator() string {
	v := c.separator
	if v == "" {
		return "_"
	}
	return v
}

// SetPrefix sets the prefix used to format environment variable names
func (c *Config) SetSeparator(separator string) *Config {
	c.separator = separator
	return c
}

func (c *Config) Prefix() string {
	v := c.prefix
	if v == "" {
		return "MYAPP"
	}
	return v
}

func (c *Config) SetPrefix(prefix string) *Config {
	c.prefix = prefix
	return c
}

// VarName formats the given name into an environment variable name
// using the rules defined by the Config.
func (c *Config) VarName(name string) string {
	return c.prefix + c.Separator() + name
}

// Lookup is the same as `os.LookupEnv`, except it formats the
// environment variable named using `VarName()` before passing
// it to `os.LookupEnv`
func (c *Config) Lookup(name string) (string, bool) {
	return os.LookupEnv(c.VarName(name))
}

// Value returns the value of the environment variable named `name`,
// except it retrieves the value using `Lookup()`. Also, if the
// value of the environment variable is empty, it returns `defaultValue`
func (c *Config) Value(name, defaultValue string) string {
	v, ok := c.Lookup(name)
	if ok {
		return v
	}
	return defaultValue
}

// SetValue sets the value of the environment variable named `name`
// to `value`, except it formats the environment variable name using
// `VarName()` before passing it to `os.Setenv`
func (c *Config) SetValue(name, value string) error {
	return os.Setenv(c.VarName(name), value)
}
