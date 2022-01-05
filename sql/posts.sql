select
    id,
    substr(title, 1, 25) title,
    substr(content, 1, 25) content,
    user_id,
    parent_id,
    strftime('%H:%M:%S', updated_at) updated_at
from
    posts
order by
    id;