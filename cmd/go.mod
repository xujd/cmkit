module cmkit

go 1.13

require (
	cmkit/pkg/auth v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.10.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/prometheus/client_golang v1.7.1
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
)

replace cmkit/pkg/auth => ../pkg/auth
