select
    id,
    title,
    content,
    user_id,
    parent_id,
    updated_at
from
    posts
order by
    updated_at desc;