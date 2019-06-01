package registry

import (
	"log"

	"github.com/go-redis/redis"

	"github.com/jinzhu/gorm"
	"github.com/speps/go-hashids"
	"github.com/tockn/diff-mvc-and-ca/external/api"
	"github.com/tockn/diff-mvc-and-ca/usecase"
)

type Registry interface {
	NewController() api.Controller
}

type registry struct {
	db      *gorm.DB
	hash    *hashids.HashID
	logger  *log.Logger
	redisDB *redis.Client
}

func NewRegistry(db *gorm.DB, hash *hashids.HashID, logger *log.Logger, red *redis.Client) Registry {
	return &registry{
		db:      db,
		hash:    hash,
		logger:  logger,
		redisDB: red,
	}
}

func (r *registry) NewController() api.Controller {
	return api.NewController(
		usecase.NewInteractor(
			api.NewPresenter(r.logger),
			r.NewItemUsecase(),
			r.NewReviewUsecase(),
		))
}
