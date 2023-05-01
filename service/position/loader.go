package position

import (
	"sync"

	"github.com/fatram/golang-dansmultipro-test/pkg/genlog"
)

var (
	positionService     *PositionService
	oncePositionService sync.Once
)

func LoadPositionService(logger genlog.Logger) *PositionService {
	oncePositionService.Do(func() {
		positionService = &PositionService{
			logger: logger,
		}
	})
	return positionService
}
