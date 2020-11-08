package req

type Request struct {
	URL string
}

// New 创建一个Request对象, 可以传入任何对象
func New(i string) *Request {
	return &Request{
		URL: i,
	}
}
