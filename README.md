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


### Git更改用户名后本地项目如何更换仓库?
```
git remote
git remote rm origin
git remote add origin https://github.com/ayunwSky/study_go.git
git push --set-upstream origin master
```
之后就可以直接用 git push 命令来推送代码了

### viper 解析多种格式的配置文件

#### viper 地址
```console
https://github.com/spf13/viper
```

#### 滴滴开源的 json 解析包
```console
https://github.com/json-iterator/go
```

#### 参考地址
```console
https://www.cnblogs.com/jiujuan/p/13799976.html
https://cloud.tencent.com/developer/article/1610531
https://zhuanlan.zhihu.com/p/105956945
```