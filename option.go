package xbox

import (
	"log"
)

type Option func(*Controller) error

func Duration(d int) Option {
	return func(c *Controller) error {
		c.duration = d
		return nil
	}
}

func Logger(l *log.Logger) Option {
	return func(c *Controller) error {
		c.logger = l
		return nil
	}
}

func AxisMargin(m int) Option {
	return func(c *Controller) error {
		c.axisMargin = m
		return nil
	}
}

func RapidFire(f bool) Option {
	return func(c *Controller) error {
		c.rapidFire = f
		return nil
	}
}
