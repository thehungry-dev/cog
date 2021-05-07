# Hungry Log

Fast, composable structured and unstructured logging for Go

## TODO

- Coloring output in handler
- Tag filtering broken for + and -

## Notes

- Logger should be a nice API for developer, has Trace, Debug, Info and similar methods
- A lower level interface of Logger is Handler
- Message should be struct to maximize performance (instead of interface)
