### Question
https://leetcode.com/problems/binary-tree-level-order-traversal/

### Accepted Status

```
[before]
34 / 34 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 3.2 MB
```

``` 
[after]
```

### before考察

* 深さごとの配列とするため、DFS探索を利用している

    LIFOである`depth`を作っているので、BFS探索でもよい。

* 隣接ノードチェックはしていない

    ノード深さとノード値がわかれば、解決できる問題だったため。

* LIFOである`depth`を使用しない方法もある？
    
    ただし、`stack`へ深さ情報を含める方法以外であるか。

### After考察
