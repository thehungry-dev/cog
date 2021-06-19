# Cog

Cog provides the building blocks for creating the logging mechanism
tailored to the need of your software.

## TODO

- Coloring output in handler
- Tag filtering broken for + and -
- Move all related stuff in separate package, keep "build" package lean to "clone and modify"

## Notes

- build should be a nice API for developer, has Trace, Debug, Info and similar methods
- A lower level interface of build is Handler
- Message should be struct to maximize performance (instead of interface)
