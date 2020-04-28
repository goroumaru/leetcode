### Question
https://leetcode.com/problems/binary-tree-level-order-traversal-ii/

### Accepted Status

```
[before]
34 / 34 test cases passed.
Status: Accepted
Runtime: 0 ms
Memory Usage: 6.5 MB
```

``` 
[after]
```

### before考察
* 出題No.102と同じ構成

    ただし、ほかの問題においてDFS探索ばかりだったので、今回はBFS探索とした。


* 深いノードから出力する

    カレントノードの深さが変わったとき、新しい配列を**先頭へ追加する**。

    そして、その配列へカレントノードを追加していく。(= 常に先頭へ追加となる)

    このとき、次のカレントノードは、BFSなので同じ階層となる。


* 上記、新しい配列を差し込むコードで、少しハマった
    
    `[][]int{}`だけで、空配列を追加した気分でいた。

    [正]

    `*levels = append([][]int{[]int{}}, *levels...) // 先頭へ追加する(空配列と結合する)`
    
    [誤]　※コンパイルは通る

    `*levels = append([][]int{}, *levels...) // 先頭へ追加する(空配列と結合する)`
    

### After考察
