package mock

// Retriever 结构体
type Retriever struct {
	Contents string
}

// Get 为结构体实现方法
//     接收者
func (r Retriever) Get(url string) string {
	return r.Contents
}
