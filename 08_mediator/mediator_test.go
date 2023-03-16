package mediator

import (
	"fmt"
	"sync"
	"testing"
)

func TestMediator(t *testing.T) {
	times := 10000
	wg := sync.WaitGroup{}
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			media := Mediator{}
			cDDriver := &CDDriver{}
			cpu := &CPU{}
			videoCard := &VideoCard{}
			soundCard := &SoundCard{}
			media.AddMember("CDDriver", cDDriver)
			media.AddMember("CPU", cpu)
			media.AddMember("VideoCard", videoCard)
			media.AddMember("SoundCard", soundCard)

			//Different goroutine input different values to verify concurrency security
			data := fmt.Sprintf("music%d,image%d", i, i)
			media.Exec("CDDriver", data)

			if cDDriver.Data != data {
				t.Fatalf("CD unexpect data %s", cDDriver.Data)
			}

			if cpu.Sound != fmt.Sprintf("music%d", i) {
				t.Fatalf("CPU unexpect sound data %s", cpu.Sound)
			}

			if cpu.Video != fmt.Sprintf("image%d", i) {
				t.Fatalf("CPU unexpect video data %s", cpu.Video)
			}

			if soundCard.Data != fmt.Sprintf("music%d", i) {
				t.Fatalf("CPU unexpect sound data %s", soundCard.Data)
			}

			if videoCard.Data != fmt.Sprintf("image%d", i) {
				t.Fatalf("CPU unexpect video data %s", videoCard.Data)
			}
		}(i)
	}
	wg.Wait()
}
