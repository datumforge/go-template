package handlers

import (
	echo "github.com/datumforge/echox"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/datumforge/datum/pkg/sessions"

	ent "github.com/datumforge/go-template/internal/ent/generated"
)

// Handler contains configuration options for handlers
type Handler struct {
	// IsTest is a flag to determine if the application is running in test mode and will mock external calls
	IsTest bool
	// DBClient to interact with the generated ent schema
	DBClient *ent.Client
	// RedisClient to interact with redis
	RedisClient *redis.Client
	// Logger provides the zap logger to do logging things from the handlers
	Logger *zap.SugaredLogger
	// ReadyChecks is a set of checkFuncs to determine if the application is "ready" upon startup
	ReadyChecks Checks
	// SessionConfig to handle sessions
	SessionConfig *sessions.SessionConfig
	// AuthMiddleware contains the middleware to be used for authenticated endpoints
	AuthMiddleware []echo.MiddlewareFunc
}
