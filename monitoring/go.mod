module github.com/gpabois/cougnat/monitoring

go 1.20

replace github.com/gpabois/cougnat/core => ../core

replace github.com/gpabois/cougnat/auth => ../auth

require (
	github.com/gpabois/cougnat/core v0.0.0-00010101000000-000000000000
	github.com/gpabois/cougnat/reporting v0.0.0-20230614205515-978629e7631c
	go.mongodb.org/mongo-driver v1.11.7
	go.uber.org/dig v1.17.0
)

require (
	github.com/PerformLine/go-stockutil v1.9.3 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jbenet/go-base58 v0.0.0-20150317085156-6237cf65f3a6 // indirect
	github.com/jdkato/prose v1.2.1 // indirect
	gopkg.in/neurosnap/sentences.v1 v1.0.6 // indirect
	k8s.io/apimachinery v0.23.4 // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gpabois/cougnat/auth v0.0.0-00010101000000-000000000000 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mongodb/mongo-tools-common v4.0.18+incompatible // indirect
	github.com/montanaflynn/stats v0.6.3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/stretchr/testify v1.8.0
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
