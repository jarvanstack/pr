# pc (Punctuation Replace)

将英文标点转化为中文标点

> 因为我写代码直接将标点符号设置为了英文, 但是在发表文章的时候英文的标点符号的展示效果不是很好, 所以写了这个工具
> 将英文标点符号转换为中文标点

## 快速开始

### 安装

```bash
$ go install github.com/zhengxiaoyao0716/pc@latest
```

或者下载编译好的文件

### 使用

```bash
$ punctuation_replace paper.txt
$ ls
paper.txt paper.txt.replace
```

例如

|  名称   | 英文  | 中文
|  ----  | ----  | ----  |
| 句号    | `. ` | `。` |
| 逗号    | `, ` | `，` |
| 分号    | `; ` | `；` |
| 冒号    | `: ` | `：` |
| 问号    | `? ` | `？` |
| 感叹号  | `! ` | `！` |
| 括号    | `(` | `（` |
| 括号    | `)` | `）` | 