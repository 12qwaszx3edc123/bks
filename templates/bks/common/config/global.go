package config

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	pbUsers "bks/proto/users"
)

var DB *gorm.DB
var Cfg Config
var RDB *redis.Client
var Ctx = context.Background()
var Esc *elastic.Client

var UsersClient pbUsers.UsersClient

