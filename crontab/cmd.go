package crontab

import "fmt"

type Cmd struct {
	spec string
	cmd  func()
}

func NewCmd(cmd func()) *Cmd {
	return &Cmd{
		cmd: cmd,
	}
}

func (c *Cmd) Cron(spec string) *Cmd {
	c.spec = spec

	return c
}

func (c *Cmd) EverySecond() *Cmd {
	return c.Cron("* * * * * *")
}

func (c *Cmd) EveryTwoSeconds() *Cmd {
	return c.Cron("*/2 * * * * *")
}

func (c *Cmd) EveryFiveSeconds() *Cmd {
	return c.Cron("*/5 * * * * *")
}

func (c *Cmd) EveryTenSeconds() *Cmd {
	return c.Cron("*/10 * * * * *")
}

func (c *Cmd) EveryFifteenSeconds() *Cmd {
	return c.Cron("*/15 * * * * *")
}

func (c *Cmd) EveryThirtySeconds() *Cmd {
	return c.Cron("*/30 * * * * *")
}

func (c *Cmd) EveryMinute() *Cmd {
	return c.Cron("0 * * * * *")
}

func (c *Cmd) EveryTwoMinutes() *Cmd {
	return c.Cron("0 */2 * * * *")
}

func (c *Cmd) EveryFiveMinutes() *Cmd {
	return c.Cron("0 */5 * * * *")
}

func (c *Cmd) EveryTenMinutes() *Cmd {
	return c.Cron("0 */10 * * * *")
}

func (c *Cmd) EveryFifteenMinutes() *Cmd {
	return c.Cron("0 */15 * * * *")
}

func (c *Cmd) EveryThirtyMinutes() *Cmd {
	return c.Cron("0 */30 * * * *")
}

func (c *Cmd) EveryHour() *Cmd {
	return c.Cron("0 0 * * * *")
}

func (c *Cmd) EveryTwoHours() *Cmd {
	return c.Cron("0 0 */2 * * *")
}

func (c *Cmd) EveryFiveHours() *Cmd {
	return c.Cron("0 0 */5 * * *")
}

func (c *Cmd) EveryTenHours() *Cmd {
	return c.Cron("0 0 */10 * * *")
}

func (c *Cmd) EveryTwelveHours() *Cmd {
	return c.Cron("0 0 */12 * * *")
}

func (c *Cmd) EveryDay() *Cmd {
	return c.Cron("0 0 0 * * *")
}

func (c *Cmd) EveryWeek() *Cmd {
	return c.Cron("0 0 0 * * 0")
}

func (c *Cmd) EveryMonth() *Cmd {
	return c.Cron("0 0 0 0 * *")
}

func (c *Cmd) EveryYear() *Cmd {
	return c.Cron("0 0 0 0 0 *")
}

func (c *Cmd) EveryWeekday() *Cmd {
	return c.Cron("0 0 0 * * 1-5")
}

func (c *Cmd) EveryWeekend() *Cmd {
	return c.Cron("0 0 0 * * 0,6")
}

func (c *Cmd) EveryDayAt(hour int) *Cmd {
	return c.Cron(fmt.Sprintf("0 0 %d * * *", hour))
}

func (c *Cmd) EveryWeekAt(week int) *Cmd {
	return c.Cron(fmt.Sprintf("0 0 0 * * %d", week))
}

func (c *Cmd) EveryMonthAt(month int) *Cmd {
	return c.Cron(fmt.Sprintf("0 0 0 %d * *", month))
}

func (c *Cmd) EveryYearAt(year int) *Cmd {
	return c.Cron(fmt.Sprintf("0 0 0 0 %d *", year))
}

func (c *Cmd) EveryWeekdayAt(hour int) *Cmd {
	return c.Cron(fmt.Sprintf("0 0 %d * * 1-5", hour))
}

func (c *Cmd) Spec() string {
	return c.spec
}

func (c *Cmd) Cmd() func() {
	return c.cmd
}
