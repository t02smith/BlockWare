# Test Data Information

## .toolkit

Used to store any toolkit data that can persist between tests.

> *For example, it is useful to store the game object of the testdir folder as that should never change*

## testdir

Represents some sample files that can be used by tests and should **always** stay the same to prevent existing tests from failing.

## tmp

A temporary folder that is is cleared before and after any tests that use it. Data put here is not meant to be permanent or used between tests.
