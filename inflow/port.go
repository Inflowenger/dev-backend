package inflow

import (
	"github.com/Inflowenger/dev-backend/env"
	fuse "github.com/Inflowenger/inflow-fusion/inflow"
)
func InitInflowConnection()error{
	return fuse.InitBackend(
		fuse.WithImplementedBackendBy(&InflowWire{}),
		fuse.WithJwtSecretKey(env.GetInfraJWTSecret()), // env INFLOW_INFRA_JWT_SECRET
		fuse.WithInfraApi(env.GetInfraApiUrl()), // env INFLOW_INFRA_API
	)

}