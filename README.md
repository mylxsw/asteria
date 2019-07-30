# Asteria 

[![Build Status](https://www.travis-ci.org/mylxsw/asteria.svg?branch=master)](https://www.travis-ci.org/mylxsw/asteria)
[![Coverage Status](https://coveralls.io/repos/github/mylxsw/asteria/badge.svg?branch=master)](https://coveralls.io/github/mylxsw/asteria?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mylxsw/asteria)](https://goreportcard.com/report/github.com/mylxsw/asteria)
[![codecov](https://codecov.io/gh/mylxsw/asteria/branch/master/graph/badge.svg)](https://codecov.io/gh/mylxsw/asteria)
[![GoDoc](https://godoc.org/github.com/mylxsw/asteria?status.svg)](https://godoc.org/github.com/mylxsw/asteria)
[![Sourcegraph](https://sourcegraph.com/github.com/mylxsw/asteria/-/badge.svg)](https://sourcegraph.com/github.com/mylxsw/asteria?badge)
[![GitHub](https://img.shields.io/github/license/mylxsw/asteria.svg)](https://github.com/mylxsw/asteria)

**Asteria** is a logging library for go.

The most straightforward way to write a log

    log.Debug("Drizzle breeze shore, dangerous night boat")
    log.Error("On the moon, the willow head, after the evening")
    log.WithFields(log.Fields{
        "user_id":  123,
        "username": "Tom",
    }).Warningf("The gentleman is frank, the villain is often jealous.")

Log according to different modules

    var logger = log.Module("asteria.user.enterprise.jobs")
       
    logger.Debug("Drizzle breeze shore, dangerous night boat")
    logger.Error("On the moon, the willow head, after the evening")
    logger.WithFields(log.Fields{
        "user_id":  123,
        "username": "Tom",
    }).Warningf("The gentleman is frank, the villain is often jealous.")
    
## Install

    go get -u github.com/mylxsw/asteria/log

## Customize

## Write the line number of the file for caller

    // default display file line number 
    log.DefaultWithFileLine(true)
    // display file line number for individual modules
    log.Module("asteria").WithFileLine(true)

### Filter

The filter supports separate settings for the specified module or global settings. With Filter, you can modify the log or cancel the log output before the log formatted output.

When multiple Filters are specified, multiple Filters are executed in the order they were added, and Global Filter takes precedence over a separate Filter set for the module.

#### Global Filter

    log.AddGlobalFilter(func(filter log.Filter) log.Filter {
        return func(f formatter.Format) {
            // if f.Level == level.Debug {
            //     return
            // }
            
            f.Fields.CustomFields["user_id"] = 123
            // Not calling filter(f) will cancel the output of the log
            filter(f)
        }
    })

#### Module Filter

    var logger = log.Module("asteria")
    logger.AddFilter(func(filter log.Filter) log.Filter {
        return func(f formatter.Format) {
            // filter(f)
            f.Level = level.Emergency
            filter(f)
        }
    })

### Log Formatter

Asteria supports custom log formats, just implement the `formatter.Formatter` interface.
    
    type Formatter interface {
        Format(f Format) string
    }

Three types of log formatting methods are provided by default

- text format, the default mode
- JSON with time
- JSON

#### Text

Use the default format, no need to make any settings, you can also specify

    // Set the default module log format
    log.Formatter(formatter.NewDefaultFormatter(true))
    // Or
    log.Default().Formatter(formatter.NewDefaultFormatter())
    // Set the log format of the specified module
    log.Module("asteria").Formatter(formatter.NewDefaultFormatter())

The format is as follows

    [RFC3339 formatted time] [logLevel] moduleName logMessage {logContext}

The module name field is specified using the `log.Module` method and the default log module is automatically generated based on the package name of the caller. Context information mainly consists of two parts

- Fields starting with `#` are automatically set by the system
- Other fields are context information set by the user using `WithFields`

Sample log output

![](https://ssl.aicode.cc/2019-07-17-15633539363228.jpg)

> You can set the default color output by `log.DefaultLogFormatter(formatter.NewDefaultFormatter(false))` or set a module to turn off color output via `log.Module("asteria").Formatter(formatter.NewDefaultFormatter(false))`.

#### Json with Time

    // Set the default module log format
    log.Formatter(formatter.NewJSONWithTimeFormatter())
    // Or
    log.Default().Formatter(formatter.NewJSONWithTimeFormatter())
    // Set the log format of the specified module
    log.Module("asteria").Formatter(formatter.NewJSONWithTimeFormatter())
 
Sample log output

    [2019-07-17T16:58:24+08:00] {"module":"user","level_name":"ERROR","level":400,"context":{"#ref":"190101931","user_id":123},"message":"user create failed","datetime":"2019-07-17T16:58:24+08:00"}
    
#### Json 

    // Set the default module log format
    log.Formatter(formatter.NewJSONFormatter())
    // Or
    log.Default().Formatter(formatter.NewJSONFormatter())
    // Set the log format of the specified module
    log.Module("asteria").Formatter(formatter.NewJSONFormatter())

Sample log output

    {"module":"asteria.user.enterprise.jobs","level_name":"EMERGENCY","level":600,"context":{"#file":"/Users/mylxsw/codes/github/asteria/log_test.go","#func":"TestModule","#line":91,"#package":"github.com/mylxsw/asteria_test","#ref":"190101931","user_id":123},"message":"He remembered the count of Monte cristo","datetime":"2019-07-17T16:58:24+08:00"}


### Log Writer

Asteria supports custom log output mode, only need to implement `writer.Writer` interface.
    
    type Writer interface {
        Write(le level.Level, module string, message string) error
        ReOpen() error
        Close() error
    }


#### Stdout

The default output mode is **standard output**, no need to make any settings, of course, you can also specify

    // Set the default module log output
    log.Writer(writer.NewStdoutWriter())
    // can also be like this
    log.Default().Writer(writer.NewStdoutWriter())
    // Set the log format of the specified module
    log.Module("asteria").Writer(writer.NewStdoutWriter())

#### File

    // Set the default module log output
    log.Writer(writer.NewSingleFileWriter("/var/log/asteria.log"))
    // can also be like this
    log.Default().Writer(writer.NewSingleFileWriter("/var/log/asteria.log"))
    // Set the log format of the specified module
    log.Module("asteria").Writer(writer.NewSingleFileWriter("/var/log/asteria.log"))
    

If you want to rotate the logs according to your own rules, such as generating new log files every day, you can use `RotatingFileWriter`

    fw := writer.NewDefaultRotatingFileWriter(func(le level.Level, module string) string {
        return fmt.Sprintf("asteria.%s.log", time.Now().Format("20060102"))
    })
    
    // you can call GC to close unused files every time you wanted
    fw.GC(1 * time.Hour)
    // Or you can call AutoGC to start auto gc support
    // asteria will call GC function every one hour for you automatically
    fw.AutoGC(context.Background(), time.Hour)
    
    log.Writer(fw)


#### Syslog

    // Set the default module log output
    log.Writer(writer.NewSyslogWriter("", "", syslog.LOG_DEBUG | syslog.LOG_SYSLOG, "asteria"))
    // can also be like this
    log.Default().Writer(writer.NewSyslogWriter("", "", syslog.LOG_DEBUG | syslog.LOG_SYSLOG, "asteria"))
    // Set the log format of the specified module
    log.Module("asteria").Writer(writer.NewSyslogWriter("", "", syslog.LOG_DEBUG | syslog.LOG_SYSLOG, "asteria"))

#### Stream

    // Set the default module log output
    log.Writer(writer.NewStreamWriter(os.Stdout))
    // can alse be like this
    log.Default().Writer(writer.NewStreamWriter(os.Stdout))
    // Set the log format of the specified module
    log.Module("asteria").Writer(writer.NewStreamWriter(os.Stdout))

#### Stack

If you want to write the log to multiple different outputs, you can use `StackWriter`

    m1 := writer.NewStdoutWriter()
    m2 := writer.NewSyslogWriter("", "", syslog.DEBUG, "asteria")
    m3 := writer.NewDefaultFileWriter("/var/log/asteria.log")

    stack := writer.NewStackWriter()
    // only write debug log to m1
    stack.PushWithLevels(m1, level.Debug)
    // only write error and emergency log to m2
    stack.PushWithLevels(m2, level.Error, level.Emergency)
    // all logs will write to m3
    stack.PushWithLevels(m3)
    
    log.Writer(stack)
    // Or
    log.Default().Writer(stack)
    // Or
    log.Module("asteria").Writer(stack)
