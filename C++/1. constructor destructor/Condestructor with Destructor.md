# [Constructot with Destructor](https://study.163.com/course/courseMain.htm?courseId=271005&_trace_c_p_k2_=9943625a68ec433ba2ef4ca75f2cf1eb)

## Constructor function

---
**How a constructor does?**

*名字与类名相同 这个函数就称之为 构造函数*

```cpp
    class X {
        int i;
        // 没有返回类型
        public:
            X();
    }
```

**当我们做了 `X x；`时 就会自动调用 `x.X();` 这是自动发生的的一个动作**

总结 ：

    构造函数时对象要创建时需要调用的

## Destructor function

当对象需要销毁的时候调用

when

    当不在它的作用域时会调用

```cpp
class Y {
    int i;
    // add a tilde(~)
    public:
        ~Y();
}
```

析构函数不能有返回值以及参数

**析构的作用**

    在类的对象当中我们可能做了某些申请资源的动作，所以我们需要在释放对象的时候做一些对之前申请的资源的释放

你一定不能不知道的 Markdown 写作技巧
