# 创建第一个项目 Hello World！

    从控制台创建 `dotnet new console [--name=first CSharp]`

## .Net5 代码
```csharp
// 所使用命名空间
using System;
// 声明命名空间
namespace firstCSharp
{
    // 定义一个类
    class Program
    {
        // 静态方法
        static void Main(string[] args) {
            // 调用 Console 下 的 WriteLine 输出字符串
            Console.WriteLine("Hello, World!");
        }
    }
}
```

## Net6 代码

```csharp
Console.WriteLine("Hello, World!");
```


*以上说的时默认生成的代码*

## 运行
    在运行以下此命令时，需要存在 .csproj 文件
    
    > 1. dotnet build

    > 2. dotnet run 

> 在 build 命令中会生成 .dll 文件 (/obj/Debug/.net版本号/ref)我们可以通过  `dotnet filename` 运行

> Eg: `dotnet firstCSharp.dll`

* C# 语句以分号结束
* 我们可以手动指定 `.csproj` 文件 ,通过参数 --project

    `dotnet run --project 路径`

## C# 库

    C# 的第三方库是在 Nuget 中进行集中管理
    我们可以通过 `dotnet restore` (需要有csproj文件)下载第三方库的引用，第三方库的引用都会在 .csproj 文件中进行集中管理

****
补充:

    编译：
    1. C# 变压器把source 文件编译成 Assenbly 
    1. Assembly 是 .NET Core 的包装与部署的单元(可以是lib或者exe(可执行文件))