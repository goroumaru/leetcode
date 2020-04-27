### Question
https://leetcode.com/problems/symmetric-tree/

### Accepted Status

```
[before]
195 / 195 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 2.9 MB
```

``` 
[after]
195 / 195 test cases passed.
Status: Accepted
Runtime: 8 ms
Memory Usage: 2.9 MB
```

### before考察
* 判定文をまとめられる
    
    最初の判定に `if nodeA == nil && nodeB == nil {` があるので、ふたつめの判定はまとめられる。

     `if nodeA == nil || nodeB == nil {`　へ変更


* Inputされたrootから再帰関数を適用できている
    
    フラクタル構造（DP：動的計画法）を考慮している


* 別の方法として、BFS探索も使える

### After考察
* 判定文をまとめた