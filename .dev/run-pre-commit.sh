#!/usr/bin/env sh

if [ -n "$SKIP_PRE_COMMIT" ]; then
  echo "✔️ Skipping pre-commit because env var SKIP_PRE_COMMIT exists and not-empty"
  exit 0
fi

CHANGED_GO_FILES=$(git diff HEAD --name-only | egrep '\.go$')

if [ -z "$CHANGED_GO_FILES" ]; then

  echo "✌️ No golang files changed ✌️"

else
  echo "🔎 Check Buildable"
  if ! make check-buildable; then
    echo "⛔ Code not buildable. Bad Code!"
    exit 1
  fi
  echo "✔️ Build OK"

  echo "🔎 Linting"
  if ! make lint; then
    echo "⛔ Code not clean, linting failed"
    exit 1
  fi
  echo "✔️ Lint OK"

  echo "🔎 Golang Imports"
  if ! make check-imports-newline; then
    echo "⛔ Found extra new lines in golang imports! That's Ugly!"
    exit 1
  fi
  echo "✔️ Golang imports OK"

  echo "🔎 Testing Coverage"
  if ! make test && make test-coverage; then
    echo "⛔ Test failed, code not robust!"
    exit 1
  fi
  echo "add converage.out in commit"
  git add coverage.out
  echo "✔️ Test Coverage OK"

fi

echo "✔️ Pre-Commit OK"
exit 0