[private]
@help:
    just --list --unsorted

# run tests
@test *args:
    go test -cover ./... {{ args }}
