package middleware

import "github.com/lukerodham/tuu"

// RequestLogger is a default/example implementation of how a logging middleware would be implemented.
func RequestLogger(next tuu.Handler) tuu.Handler {
	return func(c tuu.Context) error {
		c.Logger().WithField("url", c.Value("path")).Debug("New request")

		return next(c)
	}
}
