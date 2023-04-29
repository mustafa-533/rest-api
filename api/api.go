package api

import (
	"github.com/mustafa-533/rest-api/db"
	"go.uber.org/zap"
)

type H struct {
	db     *db.MySQL
	logger *zap.Logger
}

func NewHandler(db *db.MySQL, logger *zap.Logger) *H {
	return &H{
		db: db,
	}
}
