# Dynamic Memory Allocation

## new

>在 C 中 我们可以通过 `malloc()` 来进行空间的分配
通过 `free()` 来进行空间的释放
>但在 **C++** 中还添加了 **new** 与 **delete** 关键字来进行空间的分配与释放， 在 **C++** 更建议使用这两个关键字(运算符)

```cpp
new int;
new Stash;
//返回第一个地址
new int[10];
```

>**使用new 的时候 会分配空间，然后调用 constrcror function 进行对象的初始化工作**
>**在前面我们提到了这是运算符 所以必然会有结果 这个结果就是 所分配对象的空间的地址**
> new 会做两件事情 标记这个空间已经被分配以及这个空间的大小(以字节为单位)

## delete

> delete的时候会先调用析构函数然后在进行空间的释放
> 在对对象空间进行释放时会根据类型（在new 的时候以知道了类型，存在记录）进行析构函数的调用知道要

```cpp
delete p;
// 申请空间使用了 [] Eg: int *iterg = new int[3];
delete[] p;
```

```cpp
int num = new int[4];
int arr = new int[4];
num++;
//会存在没有记录无法释放，因为记录的是首地址
delete num;
// 只会调用首地址对应的一个析构函数，但是空间都会被全部回收
delete arr;
// 正确的释放时应该带有 []
delete[] arr
```

>存在 [] 则会告诉系统存在多个，需要多次调用析构函数

**不要使用 `delete` 释放不是 `new` 申请的空间**
对同一块空间不能释放两次
