package gologger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type CustomLogger struct {
	Out    io.Writer
	Prefix string
}

func New(out io.Writer, prefix string) CustomLogger {
	return CustomLogger{
		Out:    out,
		Prefix: prefix,
	}
}

func (c *CustomLogger) Info(message string) error {
	currTime := time.Now()
	timeFormat := currTime.Format(time.DateTime) + " "
	_, err := c.Out.Write([]byte(c.Prefix + INFO_PREFIX + timeFormat + message + "\n"))
	return err
}

func (c *CustomLogger) Debug(message string) error {
	currTime := time.Now()
	timeFormat := currTime.Format(time.DateTime) + " "
	_, err := c.Out.Write([]byte(c.Prefix + DEBUG_PREFIX + timeFormat + message + "\n"))
	return err
}

func (c *CustomLogger) Warning(message string) error {
	currTime := time.Now()
	timeFormat := currTime.Format(time.DateTime) + " "
	_, err := c.Out.Write([]byte(c.Prefix + WARNING_PREFIX + timeFormat + message + "\n"))
	return err
}

func (c *CustomLogger) Critical(message string) error {
	currTime := time.Now()
	timeFormat := currTime.Format(time.DateTime) + " "
	_, err := c.Out.Write([]byte(c.Prefix + CRITICAL_PREFIX + timeFormat + message + "\n"))
	return err
}

func (c *CustomLogger) Fatal(message string) {
	currTime := time.Now()
	timeFormat := currTime.Format(time.DateTime) + " "

	_, err := c.Out.Write([]byte(c.Prefix + FATAL_PREFIX + timeFormat + message + "\n"))

	if err != nil {
		fmt.Println(err)
	}

	os.Exit(1) // exit when fatal
}
