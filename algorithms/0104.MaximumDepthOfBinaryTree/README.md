### Question
https://leetcode.com/problems/maximum-depth-of-binary-tree/

### Accepted Status

```
[before]
39 / 39 test cases passed.
Status: Accepted
Runtime: 4 ms
Memory Usage: 5.6 MB
```

``` 
[after]
```

### before考察
* 出題No.103と構成は同じ


* DFSで全ノードを１度ずつ探索し、最大ツリー深さを求める
    
    カレントノードごと、最大深さと比較した。

* BFSでもできる

    `depth`配列にて、深さ管理しているので可能。


### After考察