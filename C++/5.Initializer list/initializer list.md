# Initializer list

```cpp
class Point {
    private:
        const float x, y;
        Point(float xa = 0.0, float ya = 0.0) : y(ya), x(xa){}
};
```

* 要先初始化 `y(ya), x(xa)` 再初始化之行构造函数
* 可以初始化任何类型的数据 可以是内置类型 也可以是自定义的类型 

```cpp
Student::Student(string s):name(s){}
initialzation
before constructor

Student::Student (string s){name = s;}
assignment
inside constructor 
string must have a default constructor
```

