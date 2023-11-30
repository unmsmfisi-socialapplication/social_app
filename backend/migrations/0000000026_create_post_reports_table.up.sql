CREATE TABLE soc_app_post_reports (
    report_id SERIAL PRIMARY KEY,
    post_id BIGINT NOT NULL,
    reported_by VARCHAR(255) NOT NULL, 
    reason VARCHAR(1000) NOT NULL,
    report_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    report_status VARCHAR(50) NOT NULL DEFAULT 'pending',
    resolved_by VARCHAR(255),
    FOREIGN KEY (post_id) REFERENCES soc_app_posts(post_id)
);
