# MSP log implementation

**V2 version only support log to LogProxy**

### Features
1. Compatible with Zap.
2. Add `package.Info()` method.
3. Support for dynamic customisation of Fields
4. Support for logging of function stack.
5. Enable logging of function stack for error level message.


By default, the log will be output to stdout and Redis.

### Usage

Check [v2/example](./v2/example) for usage.

#### Notice
Set HookConfig to `nil` to disable output to Redis.
Local develop environment should set to `nil`.

## V1

Based on logrus, log output to Redis. 
Check [example](./example) for usage.