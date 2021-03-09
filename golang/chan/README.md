## 源码文件
src/runtime/chan.go
## 数据结构
```golang
type hchan struct {
	qcount   uint           // 队列剩余元素个数
	dataqsiz uint           // 环形队列长度
	buf      unsafe.Pointer // 环形队列指针
	elemsize uint16         // 每个元素大小
	closed   uint32         // 关闭标识
	elemtype *_type         // 元素类型
	sendx    uint           // 队列下标，元素写入数据队列位置
	recvx    uint           // 队列下标，元素数据从该位置读取
	recvq    waitq          // 等待读消息的goroutine队列
	sendq    waitq          // 等待写消息的goroutine队列

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex              // 互斥锁，chan不允许并发读写
}
```