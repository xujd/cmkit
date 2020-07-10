package hello

import "fmt"

type Service interface {
	Hello(string) (string, error)
}
type HelloService struct {
}

func (s HelloService) Hello(world string) (string, error) {
	return fmt.Sprintf("Hello %s.", world), nil
}
