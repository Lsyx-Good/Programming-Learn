# Storage Allocation

## The default constructor

它没有参数，就默认为为 default constructor

```cpp
struct Y {
    float f;
    int i:
    Y(int a);
}
```

```cpp
Y y1[2] = {Y(1),Y(2)};     true

Y y2[2] = {Y(1)};   fasle
```


**因为没有默认的构造函数**
