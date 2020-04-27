### Question
https://leetcode.com/problems/same-tree/

### Accepted Status

```
[before]
57 / 57 test cases passed.
Status: Accepted
Runtime: 4 ms
Memory Usage: 6 MB
```

``` 
[after]
57 / 57 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.1 MB
```


### before考察
* DFS/BFS探索にとらわれ過ぎていて、構造が複雑になっている
    
    同じ木とみなすノード判定に注力する

* 再帰関数で捉えられている
    
    ただし、指定関数(isSameTree)を再帰関数として使えば、なお良し。


### After考察
* ２ツリーのカレントノードを比較している

* ノード遷移は、Left、Right同時とする
    
    ３ノード(Current、Left、Rightノード)を部分ツリーとして考える

* 再帰関数は、部分ツリーをひとまとまりとして考える
    
    return のLeft、Right比較までがひとまとまり。