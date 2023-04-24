// Code generated by goctl. DO NOT EDIT.

package cmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cmsPrefrenceAreaProductRelationFieldNames          = builder.RawFieldNames(&CmsPrefrenceAreaProductRelation{})
	cmsPrefrenceAreaProductRelationRows                = strings.Join(cmsPrefrenceAreaProductRelationFieldNames, ",")
	cmsPrefrenceAreaProductRelationRowsExpectAutoSet   = strings.Join(stringx.Remove(cmsPrefrenceAreaProductRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cmsPrefrenceAreaProductRelationRowsWithPlaceHolder = strings.Join(stringx.Remove(cmsPrefrenceAreaProductRelationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cmsPrefrenceAreaProductRelationModel interface {
		Insert(ctx context.Context, data *CmsPrefrenceAreaProductRelation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CmsPrefrenceAreaProductRelation, error)
		Update(ctx context.Context, data *CmsPrefrenceAreaProductRelation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsPrefrenceAreaProductRelationModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CmsPrefrenceAreaProductRelation struct {
		Id              int64         `db:"id"`
		PrefrenceAreaId sql.NullInt64 `db:"prefrence_area_id"`
		ProductId       sql.NullInt64 `db:"product_id"`
	}
)

func newCmsPrefrenceAreaProductRelationModel(conn sqlx.SqlConn) *defaultCmsPrefrenceAreaProductRelationModel {
	return &defaultCmsPrefrenceAreaProductRelationModel{
		conn:  conn,
		table: "`cms_prefrence_area_product_relation`",
	}
}

func (m *defaultCmsPrefrenceAreaProductRelationModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCmsPrefrenceAreaProductRelationModel) FindOne(ctx context.Context, id int64) (*CmsPrefrenceAreaProductRelation, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cmsPrefrenceAreaProductRelationRows, m.table)
	var resp CmsPrefrenceAreaProductRelation
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCmsPrefrenceAreaProductRelationModel) Insert(ctx context.Context, data *CmsPrefrenceAreaProductRelation) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, cmsPrefrenceAreaProductRelationRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PrefrenceAreaId, data.ProductId)
	return ret, err
}

func (m *defaultCmsPrefrenceAreaProductRelationModel) Update(ctx context.Context, data *CmsPrefrenceAreaProductRelation) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cmsPrefrenceAreaProductRelationRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PrefrenceAreaId, data.ProductId, data.Id)
	return err
}

func (m *defaultCmsPrefrenceAreaProductRelationModel) tableName() string {
	return m.table
}
