package context

type Item struct {
	// 提取字段名称
	Name string
	// 直接赋值,可选
	Value string
	// 提取字段值,一般提取获取，不进行赋值
	ExtractValue string
	// 提取字段css解析器
	CSS string
	// 提取css属性
	Attr string
	// 提取字段列表css选择器,可选参数
	BaseCSS string
	// 提取字段正则
	Do string
	// 提取字段日志
	Log string
	// 激活处理的正则表达式
	UrlReg string
}
