CREATE VIEW rankings AS
SELECT
    id as player_score_id,
    ROW_NUMBER () OVER (ORDER BY score DESC, submitted_at ASC) as rank
FROM
    player_scores