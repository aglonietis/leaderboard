CREATE VIEW rankings_monthly AS
SELECT
    id as player_score_id,
    ROW_NUMBER () OVER (ORDER BY score DESC, submitted_at ASC) as rank
FROM
    player_scores
WHERE
    submitted_at > ('now'::timestamp - '1 month'::interval)
