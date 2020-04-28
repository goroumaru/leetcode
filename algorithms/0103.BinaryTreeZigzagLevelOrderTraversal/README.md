### Question
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

### Accepted Status

```
[before]
33 / 33 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 3.1 MB
```

``` 
[after]
```

### before考察
* 出題No.102と構成は同じ

* 深さごとに出力方向を変える（出題意図のジグザク）ため、DFS探索とした
    1. 深さ偶数のとき、ノード左から並べて出力
    2. 深さ奇数のとき、ノード右から並べて出力

* BFSでもできる

    `depth`配列にて、深さ管理しているので可能。


### After考察