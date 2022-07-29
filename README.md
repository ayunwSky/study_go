# study_go

### Go 文档
```console
http://docscn.studygolang.com/doc/
中文文档: https://studygolang.com/pkgdoc
```

### make 和 new 的区别
1、make 和 new 都是用来申请内存(开辟一块内存空间)

2、new **比较少用**，一般用来给基本数据类型申请内存，如：`string`、`int`等，返回的是对应类型的指针（*string、*int）

3、make 是用于给 `slice`、`map`、`channel` 申请内存，make 函数返回的是对应的这三个类型本身。
