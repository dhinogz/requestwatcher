// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: requests.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRequest = `-- name: CreateRequest :one
INSERT INTO watcher_requests (
    method,
    ip,
    url,
    timestamp,
    accept,
    accept_encoding,
    accept_language,
    cf_connecting_ip,
    cf_ipcountry,
    cf_ray,
    cf_visitor,
    connection,
    host,
    priority,
    sec_ch_ua,
    sec_ch_ua_mobile,
    sec_ch_ua_platform,
    sec_fetch_dest,
    sec_fetch_mode,
    sec_fetch_site,
    sec_fetch_user,
    upgrade_insecure_requests,
    user_agent,
    x_forwarded_proto,
    x_real_ip,
    body,
    watcher_id
) VALUES (
    $1, -- method
    $2, -- ip
    $3, -- url
    $4, -- timestamp
    $5, -- accept
    $6, -- accept_encoding
    $7, -- accept_language
    $8, -- cf_connecting_ip
    $9, -- cf_ipcountry
    $10, -- cf_ray
    $11, -- cf_visitor
    $12, -- connection
    $13, -- host
    $14, -- priority
    $15, -- sec_ch_ua
    $16, -- sec_ch_ua_mobile
    $17, -- sec_ch_ua_platform
    $18, -- sec_fetch_dest
    $19, -- sec_fetch_mode
    $20, -- sec_fetch_site
    $21, -- sec_fetch_user
    $22, -- upgrade_insecure_requests
    $23, -- user_agent
    $24, -- x_forwarded_proto
    $25, -- x_real_ip
    $26, -- body
    $27  -- watcher_id
) RETURNING watcher_request_id, method, ip, url, timestamp, body, accept, accept_encoding, accept_language, cf_connecting_ip, cf_ipcountry, cf_ray, cf_visitor, connection, host, priority, sec_ch_ua, sec_ch_ua_mobile, sec_ch_ua_platform, sec_fetch_dest, sec_fetch_mode, sec_fetch_site, sec_fetch_user, upgrade_insecure_requests, user_agent, x_forwarded_proto, x_real_ip, watcher_id, created_at
`

type CreateRequestParams struct {
	Method                  string
	Ip                      string
	Url                     string
	Timestamp               pgtype.Timestamp
	Accept                  string
	AcceptEncoding          string
	AcceptLanguage          string
	CfConnectingIp          string
	CfIpcountry             string
	CfRay                   string
	CfVisitor               string
	Connection              string
	Host                    string
	Priority                string
	SecChUa                 string
	SecChUaMobile           string
	SecChUaPlatform         string
	SecFetchDest            string
	SecFetchMode            string
	SecFetchSite            string
	SecFetchUser            string
	UpgradeInsecureRequests string
	UserAgent               string
	XForwardedProto         string
	XRealIp                 string
	Body                    string
	WatcherID               int64
}

func (q *Queries) CreateRequest(ctx context.Context, arg CreateRequestParams) (WatcherRequest, error) {
	row := q.db.QueryRow(ctx, createRequest,
		arg.Method,
		arg.Ip,
		arg.Url,
		arg.Timestamp,
		arg.Accept,
		arg.AcceptEncoding,
		arg.AcceptLanguage,
		arg.CfConnectingIp,
		arg.CfIpcountry,
		arg.CfRay,
		arg.CfVisitor,
		arg.Connection,
		arg.Host,
		arg.Priority,
		arg.SecChUa,
		arg.SecChUaMobile,
		arg.SecChUaPlatform,
		arg.SecFetchDest,
		arg.SecFetchMode,
		arg.SecFetchSite,
		arg.SecFetchUser,
		arg.UpgradeInsecureRequests,
		arg.UserAgent,
		arg.XForwardedProto,
		arg.XRealIp,
		arg.Body,
		arg.WatcherID,
	)
	var i WatcherRequest
	err := row.Scan(
		&i.WatcherRequestID,
		&i.Method,
		&i.Ip,
		&i.Url,
		&i.Timestamp,
		&i.Body,
		&i.Accept,
		&i.AcceptEncoding,
		&i.AcceptLanguage,
		&i.CfConnectingIp,
		&i.CfIpcountry,
		&i.CfRay,
		&i.CfVisitor,
		&i.Connection,
		&i.Host,
		&i.Priority,
		&i.SecChUa,
		&i.SecChUaMobile,
		&i.SecChUaPlatform,
		&i.SecFetchDest,
		&i.SecFetchMode,
		&i.SecFetchSite,
		&i.SecFetchUser,
		&i.UpgradeInsecureRequests,
		&i.UserAgent,
		&i.XForwardedProto,
		&i.XRealIp,
		&i.WatcherID,
		&i.CreatedAt,
	)
	return i, err
}

const getWatcherRequests = `-- name: GetWatcherRequests :many
SELECT 
    wr.watcher_request_id, wr.method, wr.ip, wr.url, wr.timestamp, wr.body, wr.accept, wr.accept_encoding, wr.accept_language, wr.cf_connecting_ip, wr.cf_ipcountry, wr.cf_ray, wr.cf_visitor, wr.connection, wr.host, wr.priority, wr.sec_ch_ua, wr.sec_ch_ua_mobile, wr.sec_ch_ua_platform, wr.sec_fetch_dest, wr.sec_fetch_mode, wr.sec_fetch_site, wr.sec_fetch_user, wr.upgrade_insecure_requests, wr.user_agent, wr.x_forwarded_proto, wr.x_real_ip, wr.watcher_id, wr.created_at
FROM 
    watcher_requests wr
JOIN 
    watchers w ON w.watcher_id = wr.watcher_id
WHERE 
    w.hash_uri = $1
`

func (q *Queries) GetWatcherRequests(ctx context.Context, hashUri string) ([]WatcherRequest, error) {
	rows, err := q.db.Query(ctx, getWatcherRequests, hashUri)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WatcherRequest
	for rows.Next() {
		var i WatcherRequest
		if err := rows.Scan(
			&i.WatcherRequestID,
			&i.Method,
			&i.Ip,
			&i.Url,
			&i.Timestamp,
			&i.Body,
			&i.Accept,
			&i.AcceptEncoding,
			&i.AcceptLanguage,
			&i.CfConnectingIp,
			&i.CfIpcountry,
			&i.CfRay,
			&i.CfVisitor,
			&i.Connection,
			&i.Host,
			&i.Priority,
			&i.SecChUa,
			&i.SecChUaMobile,
			&i.SecChUaPlatform,
			&i.SecFetchDest,
			&i.SecFetchMode,
			&i.SecFetchSite,
			&i.SecFetchUser,
			&i.UpgradeInsecureRequests,
			&i.UserAgent,
			&i.XForwardedProto,
			&i.XRealIp,
			&i.WatcherID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
