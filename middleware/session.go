package middleware

import (
	"github.com/lukerodham/tuu"
)

func SessionSaver(next tuu.Handler) tuu.Handler {
	return func(c tuu.Context) error {
		err := next(c)
		if err != nil {
			return err
		}

		return c.Session().Save()
	}
}
