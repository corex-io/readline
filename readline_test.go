package readline

import (
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	rl, err := NewEx(&Config{
		Prompt: StaticPrompt(""),
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	go func() {
		for range time.Tick(time.Millisecond) {
			rl.SetPrompt(StaticPrompt("hello"))
		}
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		rl.Close()
	}()

	rl.Readline()
}
