package pkg

import (
	"article_user_service/pkg/config"
	"article_user_service/pkg/db"
	"article_user_service/pkg/logger"
	"article_user_service/pkg/migration"

	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	db.Module,
	logger.Module,
	migration.Module,
)
