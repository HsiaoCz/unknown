# 不知道干啥的项目的

yaml 格式的语法：

使用缩进表示层级关系，缩进的空格数不重要，只要相同层级的元素左对齐即可
`#`号表示注释

yaml 的基本数据类型：

纯量、数组、对象

纯量：最基本的不可再分的值，包括字符串，布尔值、整数。浮点数、null、时间、日期

```yaml
# 注释
name: redis
version: 5.6
port: 6379
stdin: true
image: null # 表示空值 还可以使用~

# date and time formart ISO 8601
date: 2022-03-18
time: 2022-03-18T08:30:10:00:00

# 如果字符串很长
siglelineString: >-
  this is  a very lang string
  anthor line xssssssss
  anthor line

# 最后解析的效果就是这样的
# this is  a very lang string anthor line xssssssss anthor line

# 想保留每一行结尾的换行符
multilineString: |
  this is  a very lang string
  anthor line xssssssss
  anthor line

# 解析的效果 this is  a very lang string/n anthor line xssssssss/n anthor line/n
```

数组：以`-`开头表示一个数组

```yaml
port:
  - 6379
  - 6380

# 等价写法
prot: [6379, 6380]
```

对象

```yaml
# 定义一个containser这样的对象：
container:
  name: mysql
  image: mysql:5.0
  prot: 1234
  version: 8.0
```

举个例子，将三个类型放在一起

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: redis-stdin

spec:
  containers:
    - name: redis
      image: redis
```

### 将请求的 json 格式的参数取出

```go
// json.NewDecoder(r).Decode(v)
dec := json.NewDecoder(r.Body)
	err := dec.Decode(&userRes)
	if err != nil {
		log.Fatal(err)
	}
```

### 使用 validator 进行参数校验

有一点需要注意，结构体的 tag 加`validate`，而不是`validator`

```go
vaildate := validator.New()
	if err := vaildate.Struct(userRes); err != nil {
		fmt.Fprintln(w, err)
		return
	}
```

**1、这留了一个问题，就是这个错误响应没有本地化，这个问题后面解决**

**2、接下来要解决的问题是，避免重复注册**
避免重复注册的功能实现

**3、接下来实现用户登录**
用户登录完成

**4、实现登录校验功能**
实现用户成功登录会拿到 token

```jwt
go get github.com/golang-jwt/jwt/v4
```

这里实现对响应信息的简单封装
通过一个 map 来封装响应消息

之前的那个问题，vaildator 参数校验，存在着返回错误信息包含敏感内容的情况
这里用了一个包装函数，将这种错误信息隐去了

**5、接下来实现中间件，错误恢复，以及日志，将日志写入文件**

**6、实现中间件校验 token，以及路由分组**
