### Question
https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

### Accepted Status

```
[before]
32 / 32 test cases passed.
Status: Accepted
Runtime: 116 ms
Memory Usage: 197.6 MB
```

``` 
[after]
```

### before考察
* すこしヒントをもらって解いた（他のEasyより難しい印象）

* Submission Detailを見ると、Runtimeにばらつきがあり、もっと高速な方法がありそう

    1. そもそも今回の方法が悪い？
    2. 今回の方法でも良いが、再帰する毎にスライスを渡してるのが悪い？

* Acceptedだけれど・・・

    Your Input `[1,2,3,4,5,6,7,8]`　→　OUTPUT `[5,3,7,2,4,6,8,1]`

    これって、あってるの？？？出題とずれてるような・・・
    
    1. そもそも出題とずれてるから、Expectedも適当な値が戻ってくる
    2. これであってる


* 部分ツリーに分けて考えた

    フラクタル構造を見つけて、再帰関数で解きたかった。    
    移動したカレントノードを部分ツリーのルートとみなす。
    カレントノードが移動する（再帰関数を呼び出す）とき、**部分ツリー全体を**渡すことで、
    元のツリーと部分ツリーがフラクタルになって、再帰関数を適用できる。


* 部分ツリーのルートインデックスの見つけ方は、ヒントをもらって。

### After考察
