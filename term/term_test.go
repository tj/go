package term

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestRenderer(t *testing.T) {
	if os.Getenv("TEST_SLOW") == "" {
		t.SkipNow()
	}

	render := Renderer()

	fmt.Printf("\n\n-----\n")
	time.Sleep(time.Second)

	render(`Deploying`)
	time.Sleep(time.Second)

	render(`Deploying
  - some resource`)
	time.Sleep(time.Second)

	render(`Deploying
  - some resource (complete)
  - another resource`)
	time.Sleep(time.Second)

	render(`Deploying
  - some resource (complete)
  - another resource (complete)`)
	time.Sleep(time.Second)

	render(`Deploying
  - some resource (complete)
  - another resource (complete)
  - final resource`)
	time.Sleep(time.Second)

	render(`Deploying
  - some resource (complete)
  - another resource (complete)
  - final resource (complete)`)
	time.Sleep(time.Second)

	render(`Deployment complete`)
	time.Sleep(time.Second * 2)

	render(``)
	time.Sleep(time.Second)

	fmt.Printf("-----\n\n\n")
}
