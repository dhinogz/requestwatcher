// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Watcher struct {
	WatcherID int64
	HashUri   string
	CreatedAt pgtype.Timestamptz
	Active    bool
}

type WatcherRequest struct {
	WatcherRequestID        int64
	Method                  string
	Ip                      string
	Url                     string
	Timestamp               pgtype.Timestamp
	Body                    string
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
	WatcherID               int64
	CreatedAt               pgtype.Timestamptz
}
