module cmkit

go 1.13

require (
	cmkit/pkg/auth v0.0.0-00010101000000-000000000000
	cmkit/pkg/fileupload v0.0.0-00010101000000-000000000000
	cmkit/pkg/hello v0.0.0-00010101000000-000000000000
	cmkit/pkg/home v0.0.0-00010101000000-000000000000
	cmkit/pkg/models v0.0.0-00010101000000-000000000000 // indirect
	cmkit/pkg/res v0.0.0-00010101000000-000000000000
	cmkit/pkg/sys v0.0.0-00010101000000-000000000000
	cmkit/pkg/utils v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/jinzhu/gorm v1.9.15
	github.com/prometheus/client_golang v1.11.1
)

replace cmkit/pkg/auth => ../pkg/auth

replace cmkit/pkg/hello => ../pkg/hello

replace cmkit/pkg/utils => ../pkg/utils

replace cmkit/pkg/models => ../pkg/models

replace cmkit/pkg/sys => ../pkg/sys

replace cmkit/pkg/res => ../pkg/res

replace cmkit/pkg/home => ../pkg/home

replace cmkit/pkg/fileupload => ../pkg/fileupload
