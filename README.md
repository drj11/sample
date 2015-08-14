# sample

```
sample [-f 0.1] [-p 10] [-n 10] [file ...]
```

Sample a file at random.

Note: only the `-f` option is supported.

## Examples

Pick one in 10 lines:

    $ sample /etc/passwd
    bin:x:2:2:bin:/bin:/usr/sbin/nologin
    sys:x:3:3:sys:/dev:/usr/sbin/nologin
    backup:x:34:34:backup:/var/backups:/usr/sbin/nologin
    messagebus:x:102:106::/var/run/dbus:/bin/false
    kernoops:x:106:65534:Kernel Oops Tracking Daemon,,,:/:/bin/false

Pick one in a million lines:

    $ yes | nl |
      sample -f 0.000001
    1105595   y
    7231072   y
    7307595   y
    8393726   y
    ^C

Same as above but more silly:

    $ yes | nl |
      sample | sample | sample | sample | sample | sample
    1258699     y
    1396821     y
    2054768     y
    2942798     y
    ^C

## Manual

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
