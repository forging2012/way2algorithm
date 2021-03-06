# 桶排序

桶排序的本质是抽屉原理，这里『桶』其实就是『抽屉』，在实现的时候就是『值域区间』。
各个桶之间是有序的，只要把桶里边的元素也弄成有序，那么拼接后整体也就有序了。

实现的时候，桶一般有两种数据结构，数组和链表，下面是其优劣：

**数组实现**:
- 优点: 桶内元素进行排序的时候，可以选择随机访问性质的排序算法，比如利用二分
    优化的插入排序。
- 缺点：由于事先不知道桶的大小，使用固定长度的数组有可能造成空间浪费，而使用
    变长数组的话可能会伴随多次allocate和数组copy。

**链表实现**：缺点和优点刚好和数组实现桶排序算法的反过来。

## 例子

下面是利用数组实现的一个例子，仔细体会其过程。

**初始状态**：
- 待排数组： array = [27 18 82 79 2 36 82 70 30 91]
- 数组最大最小元素：max=91, min=2
- 桶: buckets = [[], [], [], [], [], [], [], [], [], []]

这里我们将桶的个数设置为数组的长度`n`，则每个桶的长度bucketSize为`(max-min+1)/n`，
第`i`个桶的对应的区间为`[min+bucketSize*i, min+bucketSize*(i+1)-1]`。这样
根据元素大小，我们就能够反推出该元素所在的桶的下标。

**排序过程**：

- 第一步：分桶, 即将数组的元素放到对应的桶中。分好后桶为：

    [[2] [18] [27] [36 30] [] [] [] [70] [82 79 82] [91]]

- 第二步：桶内排序，即将桶内的元素排好序。桶内排序算法这里我们选用计数排序，排序后的桶为：

    [[2] [18] [27] [30 36] [] [] [] [70] [79 82 82] [91]]

- 第三步：填充，即把桶内的元素依次填充到原数组。
