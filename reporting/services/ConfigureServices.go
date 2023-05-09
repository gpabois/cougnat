package services

import "go.uber.org/dig"

// Configure the services
func ConfigureServices(container *dig.Container) {
	container.Provide(ProvideReportService)
}
