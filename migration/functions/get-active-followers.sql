CREATE OR REPLACE FUNCTION get_active_followers(input_person_id UUID)
    RETURNS TABLE (
        follower_id UUID,
        follower_first_name VARCHAR(32),
        follower_email_address varchar(64)
            ) AS $$
BEGIN
    RETURN QUERY
    SELECT DISTINCT f.id AS follower_id, f.first_name  AS follower_first_name, f.email_address AS follower_email_address
    FROM connection c
             JOIN person f ON c.user_id = f.id
             LEFT JOIN post p ON f.id = p.person_id
    WHERE c.connected_to = input_person_id
      AND f.status = 'Active'
      AND (
        -- Follower created a connection within the last 10 days
        EXISTS (
            SELECT 1
            FROM connection c2
            WHERE c2.user_id = f.id
              AND c2.datetime_created >= CURRENT_DATE - INTERVAL '10 days'
        )
            -- OR connection was inserted within the last 30 days
            OR f.datetime_created >= CURRENT_DATE - INTERVAL '10 days'
            -- OR connection created a post within the last 30 days
            OR EXISTS (
                SELECT 1
                FROM post p2
                WHERE p2.person_id = f.id
                  AND p2.datetime_created >= CURRENT_DATE - INTERVAL '10 days'
            )
        );
END;
$$
LANGUAGE plpgsql;