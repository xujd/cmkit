package hello

import "fmt"

// Service 服务接口
type Service interface {
	Hello(string) (string, error)
}

// HelloService 服务接口
type HelloService struct {
}

// Hello Hello
func (s HelloService) Hello(world string) (string, error) {
	return fmt.Sprintf("Hello %s.", world), nil
}
