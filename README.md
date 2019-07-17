# asteria 

[![Build Status](https://www.travis-ci.org/mylxsw/asteria.svg?branch=master)](https://www.travis-ci.org/mylxsw/asteria)
[![Coverage Status](https://coveralls.io/repos/github/mylxsw/asteria/badge.svg?branch=master)](https://coveralls.io/github/mylxsw/asteria?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mylxsw/asteria)](https://goreportcard.com/report/github.com/mylxsw/asteria)
[![codecov](https://codecov.io/gh/mylxsw/asteria/branch/master/graph/badge.svg)](https://codecov.io/gh/mylxsw/asteria)
[![GoDoc](https://godoc.org/github.com/mylxsw/asteria?status.svg)](https://godoc.org/github.com/mylxsw/asteria)

asteria is a log library for go

最简单的记录日志方式

    asteria.Debug("细雨微风岸，危樯独夜舟")
    asteria.Error("月上柳梢头，人约黄昏后")
    asteria.WithContext(asteria.C{
        "user_id":  123,
        "username": "Tom",
    }).Warningf("君子坦荡荡，小人常戚戚")

分模块记录日志

    var logger = asteria.Module("asteria.user.enterprise.jobs")
       
    logger.Debug("细雨微风岸，危樯独夜舟")
    logger.Error("月上柳梢头，人约黄昏后")
    logger.WithContext(asteria.C{
        "user_id":  123,
        "username": "Tom",
    }).Warningf("君子坦荡荡，小人常戚戚")
    
## 自定义

## 输出调用日志的文件行号

    # 设置默认显示文件行号
    asteria.DefaultWithFileLine(true)
    # 为单独模块设置显示文件行号
    asteria.Module("asteria").WithFileLine(true)

### Filter

Filter 支持单独为指定模块设置或者全局设置，通过Filter，你可以在日志格式化输出之前对日志进行统一修改或者取消日志的输出。

当指定多个 Filter 时，多个 Filter 会按照添加顺序依次执行，全局 Filter 优先于为模块设置的单独 Filter 执行。

#### 全局 Filter

    asteria.AddGlobalFilter(func(filter asteria.Filter) asteria.Filter {
		return func(f formatter.Format) {
			// if f.Level == level.Debug {
			// 	return
			// }

			f.Context.UserContext["user_id"] = 123
            
            // 不调用 filter(f) 将取消日志的输出
			filter(f)
		}
	})

#### 模块 Filter

    var logger = asteria.Module("asteria")
    logger..AddFilter(func(filter asteria.Filter) asteria.Filter {
		return func(f formatter.Format) {
			// filter(f)
			f.Level = level.Emergency
			filter(f)
		}
	})

### 日志格式

Asteria 支持自定义日志格式，只需要实现 `formatter.Formatter` 接口即可。
    
    type Formatter interface {
    	Format(f Format) string
    }

默认提供了三种类型的日志格式化方式

- 文本格式，默认方式
- Json+时间
- Json

#### 文本格式


使用默认格式，不需要进行任何设置，也可以这样指定

    // 设置默认模块日志格式
    asteria.Formatter(formatter.NewDefaultFormatter())
    // 也可以这样
    asteria.Default().Formatter(formatter.NewDefaultFormatter())
    // 设置指定模块的日志格式
    asteria.Module("asteria").Formatter(formatter.NewDefaultFormatter())

格式如下

    [RFC3339格式的时间] [日志级别] 模块名 日志内容 {上下文信息，json格式}

模块名字段在使用 `asteria.Module` 方法指定后，使用自定义的名称，默认日志模块自动根据调用日志的文件包名生成。上下文信息中主要包含两部分

- 以 `#` 开头的字段为系统自动设置的字段
- 其它字段为用户使用 `WithContext` 设置的上下文信息

日志输出样例

![](https://ssl.aicode.cc/2019-07-17-15633539363228.jpg)

> 你可以通过 `asteria.DefaultWithColor(false)` 设置默认关闭彩色输出，或者通过 `asteria.Module("asteria").WithColor(false)` 设置某个模块关闭彩色输出。

#### Json+时间格式

    // 设置默认模块日志格式
    asteria.Formatter(formatter.NewJSONWithTimeFormatter())
    // 也可以这样
    asteria.Default().Formatter(formatter.NewJSONWithTimeFormatter())
    // 设置指定模块的日志格式
    asteria.Module("asteria").Formatter(formatter.NewJSONWithTimeFormatter())
 
日志输出样例

    [2019-07-17T16:58:24+08:00] {"module":"user","level_name":"ERROR","level":400,"context":{"#ref":"190101931","user_id":123},"message":"user create failed","datetime":"2019-07-17T16:58:24+08:00"}
    
#### Json 格式

    // 设置默认模块日志格式
    asteria.Formatter(formatter.NewJSONFormatter())
    // 也可以这样
    asteria.Default().Formatter(formatter.NewJSONFormatter())
    // 设置指定模块的日志格式
    asteria.Module("asteria").Formatter(formatter.NewJSONFormatter())

日志输出样例

    {"module":"asteria.user.enterprise.jobs","level_name":"EMERGENCY","level":600,"context":{"#file":"/Users/mylxsw/codes/github/asteria/log_test.go","#func":"TestModule","#line":91,"#package":"github.com/mylxsw/asteria_test","#ref":"190101931","user_id":123},"message":"He remembered the count of Monte cristo","datetime":"2019-07-17T16:58:24+08:00"}


### 日志输出

Asteria 支持自定义日志输出方式，只需要实现 `writer.Writer` 接口即可。
    
    type Writer interface {
        Write(le level.Level, message string) error
        ReOpen() error
        Close() error
    }

默认提供了三种类型的日志输出方式

#### 标准输出

默认输出方式为 **标准输出**，不需要做任何设置，当然，也可以自己指定

    // 设置默认模块日志输出
    asteria.Writer(writer.NewStdoutWriter())
    // 也可以这样
    asteria.Default().Writer(writer.NewStdoutWriter())
    // 设置指定模块的日志格式
    asteria.Module("asteria").Writer(writer.NewStdoutWriter())

#### 文件

    // 设置默认模块日志输出
    asteria.Writer(writer.NewSingleFileWriter("/var/log/asteria.log"))
    // 也可以这样
    asteria.Default().Writer(writer.NewSingleFileWriter("/var/log/asteria.log"))
    // 设置指定模块的日志格式
    asteria.Module("asteria").Writer(writer.NewSingleFileWriter("/var/log/asteria.log"))

#### syslog

    // 设置默认模块日志输出
    asteria.Writer(writer.NewSyslogWriter("", "", syslog.LOG_DEBUG | syslog.LOG_SYSLOG, "asteria"))
    // 也可以这样
    asteria.Default().Writer(writer.NewSyslogWriter("", "", syslog.LOG_DEBUG | syslog.LOG_SYSLOG, "asteria"))
    // 设置指定模块的日志格式
    asteria.Module("asteria").Writer(writer.NewSyslogWriter("", "", syslog.LOG_DEBUG | syslog.LOG_SYSLOG, "asteria"))