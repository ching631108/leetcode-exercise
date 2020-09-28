## 该题目直接抄的答案，就不提交了

答案源自[《Leetcode Cookbook》](https://books.halfrost.com/leetcode/ChapterFour/0004.Median-of-Two-Sorted-Arrays/)

## 关于答案中一些细节的讨论

### 变量 k

解答中变量k的值如下

` k :=  (len(nums1)+len(nums2)+1)>>1 `

既两个数组长度除以2.

在确定数组1的切分线 ` nums1Mid ` 之后，使用 k 寻找 ` nums2Mid `

` nums2Mid = k - nums1Mid `

关于这件事儿我是这么理解的。既然我们要寻找的 nums1Mid 和 nums2Mid 是两个数组合并之后的中位数。那么中位数左边的数字个数是固定的。这个数字就是 k。

所以当我们找到 nums1Mid 的时候，nums2Mid 就是剩下的数字的长度