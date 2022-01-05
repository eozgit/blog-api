select
    id,
    username,
    password,
    strftime('%H:%M:%S', updated_at) updated_at
from
    users
order by
    id;