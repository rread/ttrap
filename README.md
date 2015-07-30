# Simple Trigger Trap

This is a sample program can trigger a cammera from your computer  using 
[Triggertrap kit](http://www.triggertrap.com/#products/triggertrapmobile1). This
was inspired by [Aaron Harris](http://blog.awharrisphotography.com)'s script he wrote for photographing the
International Space Station.

## Build

This has been tested on Mac and Linux. On Linux this uses `aplay` from
the `alsa-utils` package.

play the sound, 
Install [go](golang.org) and 

```
ttrap$ go build
````

## Usage

```
ttrap$ ./ttrap --at "2015-07-29 23:13:12 PDT"
2015/07/29 23:12:55 wait for it...
2015/07/29 23:13:12 click
2015/07/29 23:13:22 done
```



The tone file is derived from one on this
[post](http://www.triggertrap.com/news/call-your-phone-to-trigger/)
from TriggerTrap.
