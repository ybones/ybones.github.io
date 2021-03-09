## 源码文件
src/runtime/slice.go
## 数据结构
```golang
type slice struct {
	array unsafe.Pointer    // 指向数组地址
	len   int               // 长度
	cap   int               // 容量
}
```
## 扩容
    如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍;
    如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍;