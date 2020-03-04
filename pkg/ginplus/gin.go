package ginplus

// 定义上下文中的键
const (
	prefix = "doudou"
	// UserIDKey 存储上下文中的键(用户ID)
	UserIDKey = prefix + "/user_id"
	// TraceIDKey 存储上下文中的键(跟踪ID)
	TraceIDKey = prefix + "/trace_id"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res_body"
)
