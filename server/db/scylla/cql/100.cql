create table if not exists reports(
    id uuid,
    user_id text,

    type text,
    message text,
    created_at timestamp,
    primary key ( id )
)