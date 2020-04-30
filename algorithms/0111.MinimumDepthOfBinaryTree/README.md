### Question
https://leetcode.com/problems/minimum-depth-of-binary-tree/

### Accepted Status

```
[before]
41 / 41 test cases passed.
Status: Accepted
Runtime: 4 ms
Memory Usage: 5.5 MB
```

``` 
[after]
```

### before考察
* 問題No.110と基本構成は同じ

* 問題No.104の構成の方が、おそらく早そう

* 最小深度を求めたいので、DFS探索を使う

    No2で再帰関数から１つ抜け、抜けた先の親ノード(部分ツリーのルート)となる。
    
    No4でLeft、Rightそれぞれの再帰関数から１つ抜け、ひとつ上層のノードを探索する。

    1. 任意の親ノードのLeftを葉まで探索する（※すべてLeft）
    2. 葉ノードまで来たら、それまでの最小深さと比較する
    3. 親ノードへ戻る
    4. 1~3をRightでも実施する

### After考察
