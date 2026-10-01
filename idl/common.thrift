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

struct VueRoute {
    1: i64 id, //ID
    2: i64 parentId, //父ID
    3: string title, //导航标题
    4: i32 type, //类型 1目录 2菜单 3按钮
    5: string path, //路由地址
    6: string name, //导航名称
    7: string component, //组件路径
    8: string redirect, //路由重定向
    9: string icon, //图标
    10: string permission, //权限标识
    11: string locale, //语言包键名
    12: bool isCache, //是否缓存: 1是 0否
    13: bool isHidden, //是否隐藏: 1是 0否
    14: bool isExternal, //是否外链: 1是 0否
    15: i32 sort, //排序值
    16: string status, //状态: （1：启用；2：禁用）
    17: string activeMenu, //
    18: bool alwaysShow, //
    19: bool breadcrumb, //
    20: bool showInTabs, //
    21: bool affix, //
    22: list<string> roles,
    23: list<VueRoute> children,
}

struct Dropdown {
    1: i64 key, //ID
    2: i64 parentId, //父ID
    3: string title, //标题
    4: i32 sort, //排序
    5: list<Dropdown> children,
}

enum BoolStatus {
    TRUE = 0,
    FALSE = 1,
}