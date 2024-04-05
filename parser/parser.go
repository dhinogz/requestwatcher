package parser

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dhinogz/requestwatcher/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// Parses the requests for db.CreateRequestParams
func Request(r *http.Request, watcherID int64) (db.CreateRequestParams, error) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return db.CreateRequestParams{}, errors.New("could not read body")
	}
	defer r.Body.Close()

	crp := db.CreateRequestParams{
		Method: r.Method,
		Ip:     r.RemoteAddr,
		Url:    r.RequestURI,
		Timestamp: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		Accept:                  headerValue(r, "Accept"),
		AcceptEncoding:          headerValue(r, "Accept-Encoding"),
		AcceptLanguage:          headerValue(r, "Accept-Value"),
		CfConnectingIp:          headerValue(r, "Cf-Connecting-Ip"),
		CfIpcountry:             headerValue(r, "Cf-Ipcountry"),
		CfRay:                   headerValue(r, "Cf-Ray"),
		CfVisitor:               headerValue(r, "Cf-Visitor"),
		Connection:              headerValue(r, "Connection"),
		Host:                    headerValue(r, "Host"),
		Priority:                headerValue(r, "Priority"),
		SecChUa:                 headerValue(r, "Sec-Ch-Ua"),
		SecChUaMobile:           headerValue(r, "Sec-Ch-Ua-Mobile"),
		SecChUaPlatform:         headerValue(r, "Sec-Ch-Ua-Platform"),
		SecFetchDest:            headerValue(r, "Sec-Fetch-Dest"),
		SecFetchMode:            headerValue(r, "Sec-Fetch-Mode"),
		SecFetchSite:            headerValue(r, "Sec-Fetch-Site"),
		SecFetchUser:            headerValue(r, "Sec-Fetch-User"),
		UpgradeInsecureRequests: headerValue(r, "Upgrade-Insecure-Requests"),
		UserAgent:               headerValue(r, "User-Agent"),
		XForwardedProto:         headerValue(r, "X-Forwarded-Proto"),
		XRealIp:                 headerValue(r, "X-Real-Ip"),
		Body:                    string(body),
		WatcherID:               watcherID,
	}

	return crp, nil
}

func headerValue(r *http.Request, key string) string {
	v, _ := r.Header[key]
	return fmt.Sprintf("%+v", v)
}
