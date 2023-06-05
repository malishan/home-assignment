package health

import "github.com/malishan/home-assignment/business/health"

var healthProvider health.HealthAPIProvider

func InitHealthProvider(provider health.HealthAPIProvider) {
	healthProvider = provider
}
