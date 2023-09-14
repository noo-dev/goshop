package customerrors

import "fmt"

type ClientError struct {
	WrappedError error
}

func (b *ClientError) Error() string {
	return fmt.Sprintf("Client Error: %s", b.WrappedError.Error())
}
