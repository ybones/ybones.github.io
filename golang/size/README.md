| 类型 | 大小 |
| ------- | ------- |
|intN, uintN, floatN, complexN |	N/8个字节(例如float64是8个字节)|
|bool               |   1个字节|
|int, uint, uintptr |	1个机器字|
|*T                 |	1个机器字|
|string	            |   2个机器字(data,len)|
|[]T	            |   3个机器字(data,len,cap)|
|map	            |   1个机器字|
|func	            |   1个机器字|
|chan	            |   1个机器字|
|interface          |	2个机器字(type,value)|
|[n]T               |   n*T个机器字|