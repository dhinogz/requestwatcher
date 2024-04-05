-- +goose Up
CREATE TABLE watcher_requests (
    watcher_request_id BIGSERIAL PRIMARY KEY,

    method TEXT NOT NULL,
    ip TEXT NOT NULL,
    url TEXT NOT NULL,
    timestamp TIMESTAMP,
    body TEXT NOT NULL,
    accept TEXT NOT NULL,
    accept_encoding TEXT NOT NULL,
    accept_language TEXT NOT NULL,
    cf_connecting_ip TEXT NOT NULL,
    cf_ipcountry TEXT NOT NULL,
    cf_ray TEXT NOT NULL,
    cf_visitor TEXT NOT NULL,
    connection TEXT NOT NULL,
    host TEXT NOT NULL,
    priority TEXT NOT NULL,
    sec_ch_ua TEXT NOT NULL,
    sec_ch_ua_mobile TEXT NOT NULL,
    sec_ch_ua_platform TEXT NOT NULL,
    sec_fetch_dest TEXT NOT NULL,
    sec_fetch_mode TEXT NOT NULL,
    sec_fetch_site TEXT NOT NULL,
    sec_fetch_user TEXT NOT NULL,
    upgrade_insecure_requests TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    x_forwarded_proto TEXT NOT NULL,
    x_real_ip TEXT NOT NULL,

    watcher_id BIGSERIAL NOT NULL REFERENCES watchers (watcher_id) ON DELETE CASCADE,

    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE watcher_requests;

