package biz

import (
	"babycare/internal/biz/baby"
	"babycare/internal/biz/car"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(car.NewCarBiz,baby.NewBabyBiz)
