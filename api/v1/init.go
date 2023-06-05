package v1

import "github.com/malishan/home-assignment/business/home"

var apiProvider home.HomeAPIProvider

func InitHomeProvider(provider home.HomeAPIProvider) {
	apiProvider = provider
}
