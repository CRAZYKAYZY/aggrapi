package config

import (
	sqlc "github.com/CRAZYKAYZY/aggrapi/db/sqlc"
)

type APIConfig struct {
	DB *sqlc.Queries
}
