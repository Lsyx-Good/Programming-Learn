# C++ Access Control

## public private protected

    1. 控制类的哪一些成员可以被别人访问，哪一些只能自己访问
    2. 主要就是 自己与别人的区分
    自己: 类的成员函数  
    3. 访问权限的区分是相对于类来说的(同一个类对象之间是可以对私有的进行访问的)
    4. 这种权限的限制仅仅在于编译时刻的一个限制

## public

    公开的，所有的都可以进行访问

## private

    只有类的成员函数可以进行访问私有的成员变量或者成员函数

## protected

    自己以及子类可以进行访问

## class vs. struct

```cpp
class default to private

struct default to public 
```

在 C++ 中我们首选 `class` ，当我们所要表达的东西只是一些简单的类型的集合是可以使用 `struct`
