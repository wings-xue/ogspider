package req

type Request struct {
}

// New 创建一个Request对象, 可以传入任何对象
func New(i interface{}) *Request {
	return &Request{}
}
