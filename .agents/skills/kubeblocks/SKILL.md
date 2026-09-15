```markdown
# kubeblocks Development Patterns

> Auto-generated skill from repository analysis

## Overview
This skill teaches the core development patterns and conventions used in the `kubeblocks` repository, a Go codebase with a focus on clear structure and maintainability. You'll learn how to follow its file naming, import/export styles, commit message conventions, and testing patterns, enabling you to contribute effectively and consistently.

## Coding Conventions

### File Naming
- Use **snake_case** for all file names.
  - Example:  
    ```
    my_module.go
    some_util_test.go
    ```

### Imports
- Use **relative import paths** within the repository.
  - Example:
    ```go
    import (
        "github.com/yourorg/kubeblocks/pkg/utils"
        "./internal/helpers"
    )
    ```

### Exports
- Use **named exports** for functions, types, and variables.
  - Example:
    ```go
    // Exported function
    func DoSomething() {}

    // Exported type
    type MyStruct struct{}
    ```

### Commit Messages
- Follow **conventional commit** style.
- Prefixes: `fix`, `chore`, `feat`
- Example:
  ```
  feat: add support for custom resource definitions
  fix: correct typo in cluster reconciliation logic
  chore: update dependencies to latest minor versions
  ```

## Workflows

_No automated workflows detected in this repository._

## Testing Patterns

- Test files use the pattern `*.test.*`.
  - Example: `handler.test.go`
- The specific testing framework is not detected, but tests are written in Go.
- To write a test:
  ```go
  // handler.test.go
  package mypackage

  import "testing"

  func TestMyFunction(t *testing.T) {
      // test logic here
  }
  ```

## Commands
| Command         | Purpose                                    |
|-----------------|--------------------------------------------|
| /test           | Run all tests in the repository            |
| /lint           | Lint the codebase for style conformance    |
| /commit-guide   | Show commit message conventions            |
| /conventions    | Show coding conventions summary            |
```
