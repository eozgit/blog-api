select
    id,
    title,
    content,
    post_id,
    parent_id,
    user_id,
    deleted_at,
    updated_at
from
    comments
order by
    updated_at desc;