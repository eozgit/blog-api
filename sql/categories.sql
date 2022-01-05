select
    id,
    title,
    content,
    strftime('%H:%M:%S', updated_at) updated_at
from
    categories
order by
    id;