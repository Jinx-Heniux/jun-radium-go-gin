// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: cluster.sql

package db

import (
	"context"
	"database/sql"
)

const createCluster = `-- name: CreateCluster :one
INSERT INTO clusters (
  cluster_name,
  cluster_id,
  provider,
  k8s_version,
  url
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, cluster_name, cluster_id, provider, k8s_version, url, created_at
`

type CreateClusterParams struct {
	ClusterName string         `json:"cluster_name"`
	ClusterID   string         `json:"cluster_id"`
	Provider    sql.NullString `json:"provider"`
	K8sVersion  sql.NullString `json:"k8s_version"`
	Url         sql.NullString `json:"url"`
}

func (q *Queries) CreateCluster(ctx context.Context, arg CreateClusterParams) (Cluster, error) {
	row := q.db.QueryRowContext(ctx, createCluster,
		arg.ClusterName,
		arg.ClusterID,
		arg.Provider,
		arg.K8sVersion,
		arg.Url,
	)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.ClusterName,
		&i.ClusterID,
		&i.Provider,
		&i.K8sVersion,
		&i.Url,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCluster = `-- name: DeleteCluster :exec
DELETE FROM clusters
WHERE id = $1
`

func (q *Queries) DeleteCluster(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCluster, id)
	return err
}

const getCluster = `-- name: GetCluster :one
SELECT id, cluster_name, cluster_id, provider, k8s_version, url, created_at FROM clusters
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCluster(ctx context.Context, id int64) (Cluster, error) {
	row := q.db.QueryRowContext(ctx, getCluster, id)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.ClusterName,
		&i.ClusterID,
		&i.Provider,
		&i.K8sVersion,
		&i.Url,
		&i.CreatedAt,
	)
	return i, err
}

const listClusters = `-- name: ListClusters :many
SELECT id, cluster_name, cluster_id, provider, k8s_version, url, created_at FROM clusters
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListClustersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListClusters(ctx context.Context, arg ListClustersParams) ([]Cluster, error) {
	rows, err := q.db.QueryContext(ctx, listClusters, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cluster
	for rows.Next() {
		var i Cluster
		if err := rows.Scan(
			&i.ID,
			&i.ClusterName,
			&i.ClusterID,
			&i.Provider,
			&i.K8sVersion,
			&i.Url,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCluster = `-- name: UpdateCluster :one
UPDATE clusters
  set provider = $2,
  k8s_version = $3,
  url = $4
WHERE id = $1
RETURNING id, cluster_name, cluster_id, provider, k8s_version, url, created_at
`

type UpdateClusterParams struct {
	ID         int64          `json:"id"`
	Provider   sql.NullString `json:"provider"`
	K8sVersion sql.NullString `json:"k8s_version"`
	Url        sql.NullString `json:"url"`
}

func (q *Queries) UpdateCluster(ctx context.Context, arg UpdateClusterParams) (Cluster, error) {
	row := q.db.QueryRowContext(ctx, updateCluster,
		arg.ID,
		arg.Provider,
		arg.K8sVersion,
		arg.Url,
	)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.ClusterName,
		&i.ClusterID,
		&i.Provider,
		&i.K8sVersion,
		&i.Url,
		&i.CreatedAt,
	)
	return i, err
}
