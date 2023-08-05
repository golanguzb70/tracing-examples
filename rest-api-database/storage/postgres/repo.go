package postgres

import (
	"context"

	"github.com/golanguzb70/tracing-examples/rest-api-database/models"
)

type PostgresI interface {
	// common
	UpdateSingleField(ctx context.Context, req *models.UpdateSingleFieldReq) error
	CheckIfExists(ctx context.Context, req *models.CheckIfExistsReq) (*models.CheckIfExistsRes, error)

	// User
	UserCreate(ctx context.Context, req *models.UserCreateReq) (*models.UserResponse, error)
	UserGet(ctx context.Context, req *models.UserGetReq) (*models.UserResponse, error)
	UserFind(ctx context.Context, req *models.UserFindReq) (*models.UserFindResponse, error)
	UserUpdate(ctx context.Context, req *models.UserUpdateReq) (*models.UserResponse, error)
	UserDelete(ctx context.Context, req *models.UserDeleteReq) error

	// Product 
	ProductCreate(ctx context.Context, req *models.ProductCreateReq) (*models.ProductResponse, error)
	ProductGet(ctx context.Context, req *models.ProductGetReq) (*models.ProductResponse, error)
	ProductFind(ctx context.Context, req *models.ProductFindReq) (*models.ProductFindResponse, error)
	ProductUpdate(ctx context.Context, req *models.ProductUpdateReq) (*models.ProductResponse, error)
	ProductDelete(ctx context.Context, req *models.ProductDeleteReq) error
	
	// Don't delete this line, it is used to modify the file automatically
}
