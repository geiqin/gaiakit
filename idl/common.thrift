namespace go base


struct BaseResponse {
    1: i64 code,   // Status code, 0-success, other values-failure
    2: string msg, // Return status description
}

struct NilResponse {}

struct Pager {
    1: i32 paged,
    2: i32 total,
    3: i32 page_count,
    4: i32 page_size,
    5: i32 prev_page,
    6: i32 last_page,
}

enum BoolStatus {
    TRUE = 0,
    FALSE = 1,
}