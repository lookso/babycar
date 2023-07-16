package service

import (
	"babycare/internal/service/api"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(api.NewService)
