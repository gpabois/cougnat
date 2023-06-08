module github.com/gpabois/cougnat/monitoring

go 1.20

replace github.com/gpabois/cougnat/core => ../core

replace github.com/gpabois/cougnat/auth => ../auth

require go.uber.org/dig v1.17.0

require github.com/stretchr/testify v1.8.0 // indirect
