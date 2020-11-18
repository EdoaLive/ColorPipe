# ColorPipe
A simple yet effective terminal keywords highlighting tool. Very useful for inspecting logs, also piped in real-time.

This is also a proof of concept of how easy can be developing simple tools in go.

![yes, it's a meme](https://i.imgflip.com/4munxf.jpg)

# Usage
```
journalctl -ef | colorPipe
dmesg | colorPipe
tail -f /var/somelog.txt | colorPipe
```
Ans so on...

# TODO
* Update the readme with colorful screenshots
* Add a self-cleaning routine to avoid memory leaks
* Make it a library, so it can be used also for direct output or in other context (eg. in live browser editor?)
