package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golanguzb70/tracing-examples/rest-api-database/models"
)

func (r *postgresRepo) ProductCreate(ctx context.Context, req *models.ProductCreateReq) (*models.ProductResponse, error) {
	res := &models.ProductResponse{}
	query := r.Db.Builder.Insert("products").Columns(
		"product_name",
	).Values(req.ProductName).Suffix(
		"RETURNING id, product_name, created_at, updated_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.Id, &res.ProductName,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "ProductCreate: query.RunWith(r.Db.Db).Scan()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *postgresRepo) ProductGet(ctx context.Context, req *models.ProductGetReq) (*models.ProductResponse, error) {
	query := r.Db.Builder.Select("id, product_name, created_at, updated_at").
		From("products")

	if req.Id != 0 {
		query = query.Where(squirrel.Eq{"id": req.Id})
	} else {
		return &models.ProductResponse{}, fmt.Errorf("at least one filter should be exists")
	}
	res := &models.ProductResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.ProductName,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "ProductGet:query.RunWith(r.Db.Db).QueryRow()")
	}

	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *postgresRepo) ProductFind(ctx context.Context, req *models.ProductFindReq) (*models.ProductFindResponse, error) {
	var (
		res            = &models.ProductFindResponse{}
		whereCondition = squirrel.And{}
		orderBy        = []string{}
	)

	if strings.TrimSpace(req.Search) != "" {
		whereCondition = append(whereCondition, squirrel.ILike{"product_name": req.Search + "%"})
	}

	if req.OrderByCreatedAt != 0 {
		if req.OrderByCreatedAt > 0 {
			orderBy = append(orderBy, "created_at DESC")
		} else {
			orderBy = append(orderBy, "created_at ASC")
		}
	}

	countQuery := r.Db.Builder.Select("count(1) as count").From("products").Where("deleted_at is null").Where(whereCondition)
	err := countQuery.RunWith(r.Db.Db).QueryRow().Scan(&res.Count)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "ProductFind: countQuery.RunWith(r.Db.Db).QueryRow().Scan()")
	}

	query := r.Db.Builder.Select("id, product_name, created_at, updated_at").
		From("products").Where("deleted_at is null").Where(whereCondition)

	if len(orderBy) > 0 {
		query = query.OrderBy(strings.Join(orderBy, ", "))
	}

	query = query.Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "ProductFind: query.RunWith(r.Db.Db).Query()")
	}
	defer rows.Close()

	for rows.Next() {
		temp := &models.ProductResponse{}
		err := rows.Scan(
			&temp.Id, &temp.ProductName,
			&CreatedAt, &UpdatedAt,
		)
		if err != nil {
			return res, HandleDatabaseError(err, r.Log, "ProductFind: rows.Scan()")
		}

		temp.CreatedAt = CreatedAt.Format(time.RFC1123)
		temp.UpdatedAt = UpdatedAt.Format(time.RFC1123)
		res.Products = append(res.Products, temp)
	}

	return res, nil
}

func (r *postgresRepo) ProductUpdate(ctx context.Context, req *models.ProductUpdateReq) (*models.ProductResponse, error) {
	mp := make(map[string]interface{})
	mp["product_name"] = req.ProductName
	mp["updated_at"] = time.Now()
	query := r.Db.Builder.Update("products").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, product_name, created_at, updated_at")

	res := &models.ProductResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.Id, &res.ProductName,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "ProductUpdate: query.RunWith(r.Db.Db).QueryRow().Scan()")
	}
	res.CreatedAt = CreatedAt.Format(time.RFC1123)
	res.UpdatedAt = UpdatedAt.Format(time.RFC1123)

	return res, nil
}

func (r *postgresRepo) ProductDelete(ctx context.Context, req *models.ProductDeleteReq) error {
	query := r.Db.Builder.Delete("products").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return HandleDatabaseError(err, r.Log, "ProductDelete: query.RunWith(r.Db.Db).Exec()")
}
