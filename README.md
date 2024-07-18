### 实现了一个简易队列和三个简易的并发安全的缓存淘汰算法并进行测试
_____________________________________________________
# FIFO(First in First out)
先进先出是一种比较简单的淘汰算法。它的思想是先进先出，这是最简单、最公平的一种思想，即如果一个数据是最先进入的，那么可以认为在将来它被访问的可能性很小。达到缓存上限时，最先进入的数据会被最早淘汰掉。
_____________________________________________________
![image](https://github.com/user-attachments/assets/048a6109-1e0a-483a-ad7a-32bc422028b5)
# LRU(Least Recently Used)
LRU是一种缓存淘汰策略，当系统的缓存达到上限时优先淘汰最近最久未使用的数据，以提高数据的命中率并降低系统负载。
_____________________________________________________
![LRU](https://github.com/user-attachments/assets/48edde87-470a-44f3-8388-10d4bdab57f6)
# Sieve ([NSDI24](https://www.usenix.org/conference/nsdi24/presentation/zhang-yazhuo))
- 这是一种简单、高效、快速和可扩展的缓存淘汰算法，利用了“惰性提升”和 “快速降级” （Lazy promotion and Quick demotion）。
- Sieve的高效率来自于逐渐淘汰掉不受欢迎的数据。
- 与LRU和FIFO不同，Sieve需要多维护一个名为“手”的指针和一个额外的字段“visited”。
- 字段“visited”用于跟踪节点的状态，为“true”或者“false”。
- “手”指向候选淘汰节点，并从尾部逐渐移动到头部。
_____________________________________________________
![Sieve](https://github.com/user-attachments/assets/16b51762-0c45-49be-9354-6b3782a23ce3)
