BEGIN;

CREATE OR REPLACE PROCEDURE UNFOLLOW_AND_DELETE_SP(
    p_follower_profile_id INT,
    p_unfollowed_profile_id INT
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Insertar en la tabla SOC_APP_USER_PROFILE_UNFOLLOW
    INSERT INTO SOC_APP_USER_PROFILE_UNFOLLOW (
        follower_profile_id,
        unfollowed_profile_id,
        unfollow_date
    )
    VALUES (
        p_follower_profile_id,
        p_unfollowed_profile_id,
        NOW()
    );

    -- Eliminar de la tabla SOC_APP_USER_PROFILE_FOLLOW
    DELETE FROM SOC_APP_USER_PROFILE_FOLLOW 
    WHERE
        follower_profile_id = p_follower_profile_id
        AND following_profile_id = p_unfollowed_profile_id;

    COMMIT;
END;
$$;


COMMIT;