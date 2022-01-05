select
    id,
    substr(title, 1, 20) title,
    substr(content, 1, 20) content,
    post_id,
    parent_id,
    user_id,
    strftime('%H:%M:%S', deleted_at) deleted_at,
    strftime('%H:%M:%S', updated_at) updated_at
from
    comments
order by
    id;