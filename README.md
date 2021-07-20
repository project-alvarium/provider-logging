# provider-logging
A module facilitating basic structured logging for Alvarium related applications

## Usage
In the `pkg/factories` package there is a test file showing the ways in which messages can be logged.
Essentially, use the Write() function to log simple messages within a range of different severity levels.
Use the Error() function to log errors. These two functions write to `StdOut` and `StdErr` respectively.

In both functions, the last parameter is an optional [variadic parameter](https://blog.learngoprogramming.com/golang-variadic-funcs-how-to-patterns-369408f19085)
taking 0...X values. If you wish to supply additional arguments, provide them as key/value pairs. The examples
referred to above show this usage in order to set the Correlation-ID.

Use the `MinLogLevel` property on the `LoggingInfo` struct to set the minimum level of severity governing the logger's output.
Logged messages below this level of severity will be ignored. The order of severity is as follows, from lowest to highest.

0. Trace
1. Debug
2. Info
3. Warn
4. Error