module github.com/gpabois/cougnat/reporting

go 1.20

replace github.com/gpabois/cougnat/core => ../core

replace github.com/gpabois/cougnat/auth => ../auth

require (
	github.com/go-kit/kit v0.12.0
	github.com/gorilla/mux v1.8.0
	github.com/gpabois/cougnat/auth v0.0.0-00010101000000-000000000000
	github.com/gpabois/cougnat/core v0.0.0-00010101000000-000000000000
	github.com/jinzhu/copier v0.3.5
	github.com/stretchr/testify v1.8.0
	go.mongodb.org/mongo-driver v1.11.6
	go.uber.org/dig v1.17.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/gpabois/gostd v0.0.1 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mongodb/mongo-tools-common v4.0.18+incompatible // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/smartystreets/assertions v1.13.1 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/exp v0.0.0-20230510235704-dd950f8aeaea // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
