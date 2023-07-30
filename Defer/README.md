# Defer
In Go, the `defer` statement is used to schedule a function call to be executed later, typically just before the surrounding function returns.

The primary purpose of `defer` is to ensure that certain cleanup or finalization tasks are performed regardless of how the function exits (e.g., through a return statement, panic, or reaching the end of the function).

When a `defer` statement is encountered, the functionCall and its arguments are evaluated immediately, but the actual execution of the function is deferred until the surrounding function returns. 

> Multiple defer statements can be used within the same function, and they are executed in reverse order (last in, first out) during the function

