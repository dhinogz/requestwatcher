-- name: CreateRequest :one
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
) RETURNING *;

-- name: GetWatcherRequests :many
SELECT 
    wr.*
FROM 
    watcher_requests wr
JOIN 
    watchers w ON w.watcher_id = wr.watcher_id
WHERE 
    w.hash_uri = $1;
