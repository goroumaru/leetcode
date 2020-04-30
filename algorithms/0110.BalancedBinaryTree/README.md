### Question
https://leetcode.com/problems/balanced-binary-tree/

### Accepted Status

```
[before]
227 / 227 test cases passed.
Status: Accepted
Runtime: 8 ms
Memory Usage: 5.8 MB
```

``` 
[after]
```

### before考察
* 部分ツリーごとに、バランスとれているかチェックする

    1. カレントノードごと、そのノードを部分ツリーのルートとみなす
    2. ルートのLeft/Rightノードの**葉までの**深さを比較する
    3. 深度差が２以上であれば、NGとする


* Left/Rightの深度差を求めたいので、DFS探索を使う

    ***末端ノードから考えると、理解しやすい。***

    No2、No4で再帰関数から１つずつ抜け、抜けた先の親ノード(部分ツリーのルート)となる。
    
    No6で再帰関数から１つ抜け、ひとつ上層の部分ツリーバランスを確認していく。


    1. 任意ノードのLeftを葉まで探索する（※すべてLeft）
    2. 末端ノードまで来たら、親ノードへ戻る
    3. 任意ノードのRightを葉まで探索する（※すべてRight）
    4. 末端ノードまで来たら、親ノードへ戻る
    5. 親ノードにおいて、Left/Rightの深さを比べる
    6. この部分ツリーの最大深さを次の親ノードへ渡す

* Submission Detailを見てみると、Runtimeを半分くらいにできそう

    1. 部分ツリーのバランス判定結果(`leftOk` or `rightOk`)時点で、抜ける。
    2. そもそも、他の方法でないと無理？

### After考察
