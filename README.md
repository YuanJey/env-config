# env-config

`env-config` 是一个轻量级的 Go 语言库，用于将环境变量加载到结构体中。

## 特性

- **简单易用**：通过 struct tag 定义配置。
- **类型支持**：支持 `string`, `int`, `bool`, `[]string` 等基本类型。
- **默认值**：支持通过 `def` 标签指定默认值。
- **嵌套结构**：支持递归加载嵌套的结构体配置。
- **错误处理**：提供清晰的错误返回。

## 安装

```bash
go get github.com/YuanJey/env-config
```

## 使用示例

```go
package main

import (
	"fmt"
	"log"
	
	"github.com/YuanJey/env-config/pkg/load"
)

type Config struct {
	Host  string   `env:"HOST" def:"localhost"`
	Port  int      `env:"PORT" def:"8080"`
	Debug bool     `env:"DEBUG" def:"true"`
	Addrs []string `env:"ADDRS" def:"127.0.0.1"`
}

func main() {
	var cfg Config
	
	// 加载环境变量到结构体
	if err := load.LoadEnv(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Config: %+v\n", cfg)
}
```
