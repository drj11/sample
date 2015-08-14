# sample

```
sample [-f 0.1] [-p 10] [-n 10] [file ...]
```

Sample a file at random.

Note: only the `-f` option is supported.

In fraction mode (the default, or when `-f` or `-p` is used)
each line of the input is copied to the output with probability
*f*. Thus *f* is the expected fraction of the input lines that
are preserved in the output. The default value for *f* is 0.1
(10%), it can be set with the `-f` option. If the `-p n` option is
used then *f* is set to 1/n (the value of the `-p` option is the
average period between lines being copied, thus `-p 100` means
that roughly 1 in 100 lines are copied).

In count mode (when `-n` is used) the output has a fixed number
of lines.
