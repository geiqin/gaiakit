namespace go base


struct BaseResponse {
    1: i64 code,   // Status code, 0-success, other values-failure
    2: string msg, // Return status description
}

struct NilResponse {}

struct Pager {
    1: i64 paged,
    2: i64 total,
    3: i64 page_count,
    4: i64 page_size,
    5: i64 prev_page,
    6: i64 last_page,
}

enum BoolStatus {
    TRUE = 0,
    FALSE = 1,
}