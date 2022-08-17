### flag命令行参数解析

#### 1. flag.Type测试
```
go build

# 保持默认
$ ./useFlag.exe 
name: 默认名字
age: 24
married: false

# 命令行传参
$ ./useFlag.exe -name ayunwSky -age 26 -married=true
name: ayunwSky
age: 26
married: true 
```

#### 2. flag.TypeVar 测试
```
go build

# 保持默认
$ ./useFlag.exe 
姓名: 默认名字
年龄: 18
是否已婚？: false
体重: 60

# 命令行传参
$ ./useFlag.exe -name ayunwSky -age 26 -married=true -w=60.5
姓名: ayunwSky
年龄: 26
是否已婚？: true
体重: 60.5
```

#### 3. flag.Var 测试
```
$ go build
$ ./useFlag.exe -likes=篮球,足球,游戏
[篮球 足球 游戏]
```

#### 4. flag.Args 测试
```
$ go build

$ ./useFlag.exe 张三 18 男
[张三 18 男]

$ ./useFlag.exe -name=张三 -age=18 -sex=男
flag provided but not defined: -name
Usage of E:\gopath\src\github.com\useFlag\useFlag.exe:
```

#### 5. flag.Arg(i) 测试
```
$ go build

$ ./useFlag.exe 张三 18 男
索引=0，v=张三
索引=1，v=18
索引=2，v=男

$ ./useFlag.exe 张三 男 18
索引=0，v=张三
索引=1，v=男
索引=2，v=18

$ ./useFlag.exe 18  男 张三
索引=0，v=18
索引=1，v=男
索引=2，v=张三
```

参考：
```console
http://liuqh.icu/2021/06/01/go/package/9-flag/
```
