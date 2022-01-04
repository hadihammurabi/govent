package main

import (
	"errors"
	"fmt"

	"github.com/hadihammurabi/govent"
	"github.com/hadihammurabi/govent/option"
)

func main() {

	// run all listeners even if there is listener return error
	busWithSkipOnError := govent.New(option.SkipOnError)
	busWithSkipOnError.On("click", func(e govent.Event) error {
		fmt.Println("[busWithSkipOnError] clicked 1", e.Args)
		return errors.New("waduh")
	})

	busWithSkipOnError.On("click", func(e govent.Event) error {
		fmt.Println("[busWithSkipOnError] clicked 2", e.Args)
		return nil
	})
	busWithSkipOnError.Emit("click", "mantap", "asiap")

	// only execute last listener
	busWithSingleListener := govent.New(option.SingleListener)
	busWithSingleListener.On("click", func(e govent.Event) error {
		fmt.Println("[busWithSingleListener] clicked 1", e.Args)
		return nil
	})

	busWithSingleListener.On("click", func(e govent.Event) error {
		fmt.Println("[busWithSingleListener] clicked 2", e.Args)
		return nil
	})
	busWithSingleListener.Emit("click", "mantap", "asiap")
}
