package service

import (
	"babycare/internal/service/baby"
	"babycare/internal/service/car"
	"babycare/internal/service/tree"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(car.NewCarService, baby.NewBabyService, tree.NewTreeService)
