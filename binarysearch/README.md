# Binary Search

## 应用场合

**序列可以根据特定条件被分为几个部分，求某两个部分之间的临界点。**

请根据下面例子仔细体会上面这句话的含义。

### 例1

```text
给一个有序数列list，求某一个数num是否在该序列中。比如list = [1, 2, 7, 10, 12, 17, 20], num=10。
```

在这个例子中，`特定条件`为：list中的数字和num之间的大小。具体表现为：list中所有
比num小的数都在num的左边，所有比num大的数都在num的右边，跟num相等的数字在中间。

该例子可以有很多变形：
- 求num在数组中的下标
- 求序列中第一个 `>= num` 的数及其下标
- 求序列中最后一个 `< num` 的数及其下标
- 等等...

### 例2

```text
一组小球排成一排，小球可以被涂上红、黄、蓝三种颜色，相同颜色的小球排在一起，且红色的在左边，
黄色的在中间，蓝色的在右边。现在给这么一排小球，求黄色小球的个数，要求算法复杂度为O(log(n)).
```

在这个例子中，`特定条件`为: 小球的颜色。解决方案为：找出黄色小球的左右边界，然后相减得到长度。