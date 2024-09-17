#### free list

- 由于我们的b树是不可变的，因此对KV存储的每次更新都会在路径中创建新的节点，而不是更新当前的节点，这使得一些节点无法从最新版本中进行访问。
- 我们需要从旧版本中重用这些不可访问的节点，否则，数据库文件将无限增长
- 为了重用这些页面，我们将添加一个持久的免费列表来跟踪未使用的页面。更新操作在添加新页面之前更新列表中的页面，并将当前版本中未使用的页面添加到列表中。
- 该列表用作堆栈（首次退出），每个更新操作既可以从列表中删除，也可以添加到列表的顶部
- 免费列表也像我们的b树一样是不可变的。每个节点包含：
  - 1.指向未使用的页面的多个指针。
  - 2.连接到下一个节点的链接。
  - 3.列表中的项目总数。这只适用于头节点。
```shell
 | node1 | | node2 | | node3 |
  +-----------+     +-----------+     +-----------+
  | total=xxx | | | | |
  | next=yyy | ==> | next=qqq | ==> | next=eee | ==> ...
  | size=zzz | | size=ppp | | size=rrr |
  |指针||指针||指针|
  节点格式：
  |类型|大小|总|下一个|指针|
  | 2B | 2B | 8B | 8B |大小*8B  |
```

#### 用于访问列表节点的函数
```shell
func flnSize(node BNode) int
func flnNext(node BNode) uint64
func flnPtr(node BNode, idx int)
func flnSetPtr(node BNode, idx int, ptr uint64)
func flnSetHeader(node BNode, size uint16, next uint64)
func flnSetTotal(node BNode, total uint64)
```