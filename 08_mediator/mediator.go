package mediator

import (
	"fmt"
	"strings"
)

type IMediator interface {
	AddMember(string, IMember)
	Exec(string, int)
}

type IMember interface {
	Process(string)
}

type Mediator struct {
	MemMap map[string]IMember
}

func (mediator *Mediator) AddMember(key string, member IMember) {
	if mediator.MemMap == nil {
		mediator.MemMap = make(map[string]IMember)
	}
	mediator.MemMap[key] = member
}

func (mediator *Mediator) Exec(key, data string) {
	switch i := mediator.MemMap[key].(type) {
	case *CDDriver:
		i.Process(data)
		mediator.Exec("CPU", data)
	case *CPU:
		i.Process(data)
		mediator.Exec("VideoCard", i.Video)
		mediator.Exec("SoundCard", i.Sound)
	case *VideoCard:
		i.Process(data)
	case *SoundCard:
		i.Process(data)
	}
}

// Members

type CDDriver struct {
	Data string
}

func (c *CDDriver) Process(data string) {
	c.Data = data

	fmt.Printf("CDDriver: reading data %s\n", c.Data)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Process(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Process(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
}
