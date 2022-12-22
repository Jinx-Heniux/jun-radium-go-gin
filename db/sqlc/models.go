// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
)

type Application struct {
	ID         int64          `json:"id"`
	Provider   sql.NullString `json:"provider"`
	Product    sql.NullString `json:"product"`
	AppName    string         `json:"app_name"`
	AppVersion string         `json:"app_version"`
	CreatedAt  sql.NullTime   `json:"created_at"`
}

type Cluster struct {
	ID          int64          `json:"id"`
	ClusterName string         `json:"cluster_name"`
	ClusterID   string         `json:"cluster_id"`
	Provider    sql.NullString `json:"provider"`
	K8sVersion  sql.NullString `json:"k8s_version"`
	Url         sql.NullString `json:"url"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type Module struct {
	ID            int64          `json:"id"`
	ModuleName    sql.NullString `json:"module_name"`
	ModuleVersion sql.NullString `json:"module_version"`
	AppID         sql.NullInt64  `json:"app_id"`
	CreatedAt     sql.NullTime   `json:"created_at"`
}

type Project struct {
	ID          int64          `json:"id"`
	ProjectName string         `json:"project_name"`
	ProjectID   string         `json:"project_id"`
	ClusterID   string         `json:"cluster_id"`
	Url         sql.NullString `json:"url"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}
