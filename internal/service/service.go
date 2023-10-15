package service

import (
	"babycare/internal/service/baby"
	"babycare/internal/service/car"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(car.NewCarService,baby.NewBabyService)
