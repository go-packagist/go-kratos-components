package crontab

type Schedule struct {
	cmds []*Cmd
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (s *Schedule) Call(cmd func()) *Cmd {
	c := NewCmd(cmd)

	s.cmds = append(s.cmds, c)

	return c
}

func (s *Schedule) Cmds() []*Cmd {
	return s.cmds
}
