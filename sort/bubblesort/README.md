# 冒泡排序

冒泡排序过程就像其名字一样。从前往后遍历一趟，如果当前元素比其紧挨着的后一个
元素大，那么就交换两个元素的位置，总共需要遍历n - 1趟。

下面是一个例子，仔细观察，体会其过程。

    33 30 64 90 85 63 68 9 93 7
    ↑_________________________↑

    30 33 64 85 63 68 9 90 7 93
    ↑______________________↑

    30 33 64 63 68 9 85 7 90 93
    ↑___________________↑

    30 33 63 64 9 68 7 85 90 93
    ↑________________↑

    30 33 63 9 64 7 68 85 90 93
    ↑_____________↑

    30 33 9 63 7 64 68 85 90 93
    ↑__________↑

    30 9 33 7 63 64 68 85 90 93
    ↑_______↑

    9 30 7 33 63 64 68 85 90 93
    ↑____↑

    9 7 30 33 63 64 68 85 90 93
    ↑_↑

    7 9 30 33 63 64 68 85 90 93
