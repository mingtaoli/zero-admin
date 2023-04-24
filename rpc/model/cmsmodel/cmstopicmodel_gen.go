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
	cmsTopicFieldNames          = builder.RawFieldNames(&CmsTopic{})
	cmsTopicRows                = strings.Join(cmsTopicFieldNames, ",")
	cmsTopicRowsExpectAutoSet   = strings.Join(stringx.Remove(cmsTopicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cmsTopicRowsWithPlaceHolder = strings.Join(stringx.Remove(cmsTopicFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cmsTopicModel interface {
		Insert(ctx context.Context, data *CmsTopic) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CmsTopic, error)
		Update(ctx context.Context, data *CmsTopic) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsTopicModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CmsTopic struct {
		Id             int64          `db:"id"`
		CategoryId     sql.NullInt64  `db:"category_id"`
		Name           sql.NullString `db:"name"`
		CreateTime     sql.NullTime   `db:"create_time"`
		StartTime      sql.NullTime   `db:"start_time"`
		EndTime        sql.NullTime   `db:"end_time"`
		AttendCount    sql.NullInt64  `db:"attend_count"`    // 参与人数
		AttentionCount sql.NullInt64  `db:"attention_count"` // 关注人数
		ReadCount      sql.NullInt64  `db:"read_count"`
		AwardName      sql.NullString `db:"award_name"`  // 奖品名称
		AttendType     sql.NullString `db:"attend_type"` // 参与方式
		Content        sql.NullString `db:"content"`     // 话题内容
	}
)

func newCmsTopicModel(conn sqlx.SqlConn) *defaultCmsTopicModel {
	return &defaultCmsTopicModel{
		conn:  conn,
		table: "`cms_topic`",
	}
}

func (m *defaultCmsTopicModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCmsTopicModel) FindOne(ctx context.Context, id int64) (*CmsTopic, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cmsTopicRows, m.table)
	var resp CmsTopic
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

func (m *defaultCmsTopicModel) Insert(ctx context.Context, data *CmsTopic) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, cmsTopicRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CategoryId, data.Name, data.StartTime, data.EndTime, data.AttendCount, data.AttentionCount, data.ReadCount, data.AwardName, data.AttendType, data.Content)
	return ret, err
}

func (m *defaultCmsTopicModel) Update(ctx context.Context, data *CmsTopic) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cmsTopicRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CategoryId, data.Name, data.StartTime, data.EndTime, data.AttendCount, data.AttentionCount, data.ReadCount, data.AwardName, data.AttendType, data.Content, data.Id)
	return err
}

func (m *defaultCmsTopicModel) tableName() string {
	return m.table
}
