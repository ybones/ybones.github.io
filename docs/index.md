# golang源码
## chan
### 源码文件
src/runtime/chan.go
### 数据结构
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
## map
### 源码文件
src/runtime/map.go
### 数据结构
```golang
// A header for a Go map.
type hmap struct {
	count     int       // 键值对个数
	flags     uint8     // 状态标识
	B         uint8     // bucket数组的大小
	noverflow uint16 
	hash0     uint32    // hash 因子

	buckets    unsafe.Pointer // bucket数组指针，数组的大小为2^B
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

// mapextra holds fields that are not present on all maps.
type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	nextOverflow *bmap
}

// A bucket for a Go map.
type bmap struct {
    tophash [bucketCnt]uint8
}
// 编译期重构的结构 bmap
type bmap struct {
    tophash  [8]uint8       // 每个key哈希值的高8位，加速访问；
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
```
![](../golang/map/assets/bmap.png)
### 注解
 - map使用hash表实现，并使用拉链法解决冲突。
 - map的value不可寻址,扩容过程中会发生键值对迁移。
 - map遍历时加入了随机数，让每次遍历的起始bucket的位置不一样。
 - 一个hmap存放了一个（数组大小为2^B个）buckets数组指针，
每个buckets可以存储8个键值对，当每个bucket存储的kv对到达8个之后，
会通过overflow指针指向一个新的bucket，从而形成一个链表。
随着键值对数量的增加，溢出桶的数量和哈希的装载因子也会逐渐升高，超过一定范围就会触发扩容，扩容会将桶的数量翻倍。

### key 定位过程
    key 经过哈希计算后得到哈希值，共 64 个 bit 位（64位机），
    计算它到底要落在哪个桶时，只会用到最后 B 个 bit 位。
    如果 B = 5，那么桶的数量，也就是 buckets 数组的长度是 2^5 = 32。
    eg : 10010111 | 000011110110110010001111001010100010010110010101010 │ 01010
    高八位 10010111 为桶的位置
    低B（5）位 01010 为桶的编号

### 插入
    装载因子已经超过 6.5；翻倍扩容
    // Maximum average load of a bucket that triggers growth is 6.5.
    // Represent as loadFactorNum/loadFactDen, to allow integer math.
    loadFactorNum = 13
    loadFactorDen = 2
    
 - 装载因子:=元素数量÷桶数量；
 - 哈希在存储元素过多时会触发扩容操作，每次都会将桶的数量翻倍，
 扩容过程不是原子的，而是通过 runtime.growWork 增量触发的，
 在扩容期间访问哈希表时会使用旧桶，向哈希表写入数据时会触发旧桶元素的分流(取模或者位操作来获取桶的编号)。
 除了这种正常的扩容之外，为了解决大量写入、删除造成的内存泄漏问题，哈希引入了 sameSizeGrow 这一机制，在出现较多溢出桶时会整理哈希的内存减少空间的占用。

### 删除
    哈希使用了太多溢出桶；等量扩容

### 查找
    1. 根据key计算出哈希值
    2. 根据哈希值低位确定所在bucket
    3. 根据哈希值高8位确定在bucket中的存储位置
    4. 当前bucket未找到则查找对应的overflow bucket。
    5. 对应位置有数据则对比完整的哈希值，确定是否是要查找的数据
    6. 如果当前处于map进行了扩容，处于数据搬移状态，则优先从oldbuckets查找。

## slice
### 源码文件
src/runtime/slice.go
### 数据结构
```golang
type slice struct {
	array unsafe.Pointer    // 指向数组地址
	len   int               // 长度
	cap   int               // 容量
}
```
### 扩容
    如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍;
    如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍;

## unsafe.Sizeof
![](../golang/size/assets/sizeof.png)

# MySQL
## 事务 ACID
- `原子性`(Atomicity)：一个事务（transaction）中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节。事务在执行过程中发生错误，会被回滚（Rollback）到事务开始前的状态，就像这个事务从来没有执行过一样。
- `一致性`(Consistency)：在事务开始之前和事务结束以后，数据库的完整性没有被破坏。这表示写入的资料必须完全符合所有的预设规则，这包含资料的精确度、串联性以及后续数据库可以自发性地完成预定的工作。
- `隔离性`(Isolation)：数据库允许多个并发事务同时对其数据进行读写和修改的能力，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据的不一致。事务隔离分为不同级别，包括读未提交（Read uncommitted）、读提交（read committed）、可重复读（repeatable read）和串行化（Serializable）。
- `持久性`(Durability)：事务处理结束后，对数据的修改就是永久的，即便系统故障也不会丢失。
## 隔离级别
## 锁
## 索引
## 存储引擎
# Redis
## 基本数据结构及其实现
### string
### hash
### list
### set
### zset
## 缓存穿透
## 缓存击穿
## 缓存雪崩
## 分布式锁

# leetcode
- [两数之和](../leetcode/两数之和/readme.md)
- [两数相加](../leetcode/两数相加/readme.md)
- [具有所有最深节点的最小子树](../leetcode/具有所有最深节点的最小子树/readme.md)
- [合并两个有序链表](../leetcode/合并两个有序链表/readme.md)
- [回文数](../leetcode/回文数/readme.md)
- [整数反转](../leetcode/整数反转/readme.md)
- [无重复字符的最长子串](../leetcode/无重复字符的最长子串/readme.md)
- [最长公共前缀](../leetcode/最长公共前缀/readme.md)
- [有效的括号](../leetcode/有效的括号/readme.md)
- [栈的最小值](../leetcode/栈的最小值/readme.md)
- [移除元素](../leetcode/移除元素/readme.md)
- [罗马数字转整数](../leetcode/罗马数字转整数/readme.md)

# 互斥锁
## 互斥锁的实现机制
互斥锁是并发控制的一个基本手段，是为了避免竞争而建立的一种并发控制机制。

在并发编程中，如果程序中的一部分会被并发访问或修改，那么，为了避免并发访问导致的意想不到的结果，这部分程序需要被保护起来，这部分被保护起来的程序，就叫做`临界区`。

可以说，临界区就是一个被共享的资源，或者说是一个整体的一组共享资源，比如对数据库的访问、对某一个共享数据结构的操作、对一个 I/O 设备的使用、对一个连接池中的连接的调用，等等。

如果很多线程同步访问临界区，就会造成访问或操作错误，这当然不是我们希望看到的结果。所以，我们可以使用互斥锁，限定临界区只能同时由一个线程持有。
当临界区由一个线程持有的时候，其它线程如果想进入这个临界区，就会返回失败，或者是等待。直到持有的线程退出临界区，这些等待线程中的某一个才有机会接着持有这个临界区。
## 同步原语适用场景
- 共享资源。并发地读写共享资源，会出现数据竞争（data race）的问题，所以需要 Mutex、RWMutex 这样的并发原语来保护。
- 任务编排。需要 goroutine 按照一定的规律执行，而 goroutine 之间有相互等待或者依赖的顺序关系，我们常常使用 WaitGroup 或者 Channel 来实现。
- 消息传递。信息交流以及不同的 goroutine 之间的线程安全的数据交流，常常使用 Channel 来实现。