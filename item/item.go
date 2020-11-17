package item

import (
	"reflect"
)

type Field struct {
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
	// 入口URL
	StartURL []string
	// 指明下载器
	// 默认rod
	Download string
	// Active   bool
}

func HasAttr(name string, reverse bool) func(f Field) bool {
	return func(f Field) bool {
		value := reflect.ValueOf(f).FieldByName(name)
		out := value.String()
		if reverse {
			return out == ""
		}
		return out != ""
	}
}

// 过滤
func Filter(field []*Field, f func(f Field) bool) []*Field {
	var out []*Field
	for _, x := range field {
		if f(*x) {
			out = append(out, x)
		}
	}
	return out
}

func Append(field1 [][]*Field, field2 []*Field) [][]*Field {
	out := make([][]*Field, 0)
	for _, f := range field1 {
		f = append(f, field2...)
		out = append(out, f)
		// out = append(f, field2...)
	}
	return out
}

func FindKey(key string, field []*Field) *Field {
	for _, f := range field {
		if f.Name == key {
			return f
		}
	}
	return &Field{}
}

// func FindReq(field []*Field) []*Field {
// 	out := make([]*Field, 0)
// 	if len(field) == 0 {
// 		return out
// 	}
// 	for _, each := range field {
// 		if each.Active {
// 			out = append(out, each)
// 		}
// 	}
// 	return out
// }
