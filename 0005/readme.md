## 先说回文字符串的事情，和解题无关

1. 我第一个思路是，如果一个字符串，调转过来之后（从左向右读改为从右向左读），字符串内容和原本一致，则可以认为该字符串为回文字符串

2. 然后我尝试将字符串分为前半部分和后半部分，如果只调转后半部分，则调转后的字符串应该和前半部分重合（仿佛计算量少了一半）

3. 但是上述规律并不能继续拆分

## 关于本题的思路

- 先思考最短回文字符串（先剔除掉一个字符的情况），应该是两个字符组成或者是三个字符组成。AA或者是ABA的形式

- 任何回文字符串都可以由最小回文字符串向两侧扩张形成

- 我认为应该存在更好的计算方式，这个思路对于处理“aaaaaaa”这种字符串效率极低