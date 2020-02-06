package sadpath_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/sadpath"
	"github.com/stretchr/testify/assert"
)

func TestSadpathHandleCatchesCheckFailure(t *testing.T) {
	caught := false
	theError := errors.New("some error")

	func() {
		defer sadpath.Handle(func(err error) {
			caught = true
			assert.Equal(t, theError, err)
		})

		sadpath.Check(theError)
	}()

	assert.True(t, caught)
}

func TestSadpathCheckDoesNotThrowIfNoError(t *testing.T) {
	caught := false

	func() {
		defer sadpath.Handle(func(err error) {
			caught = true
		})

		sadpath.Check(nil)
	}()

	assert.False(t, caught)
}

func TestSadpathHandleRepanicsOnOtherPanic(t *testing.T) {
	caught := false
	panicVal := "boo"

	assert.PanicsWithValue(t, panicVal,
		func() {
			defer sadpath.Handle(func(err error) {
				caught = true
			})

			panic(panicVal)
		},
	)

	assert.False(t, caught)
}

func Example() {
	defer sadpath.Handle(func(err error) {
		fmt.Println(err)
	})

	var err error
	sadpath.Check(err)

	err = errors.New("oops")
	sadpath.Check(err)

	fmt.Println("never reached")

	// Output:
	// oops
}
