# Emogrid

Emogrid produces a macros to be pasted into Slack to produce emoji grids.

## Usage

```sh
emogrid [flags] <prefix> [x] [y]
Where x is the number of emojis in the line (default:  2)
      y is the number of lines              (default:  2)

Flags:  -start int
    	Start with this number (default 1)
```

For example, running:
```sh
emogrid topper 2 2
```
produces:
```
:topper-1::topper-2:
:topper-3::topper-4:
```

On a MacOS, it's easy to copy that directly into your clipboard:
```sh
emogrid topper | pbcopy
```
I'm sure that on Linux and Windows there's also a command like that, but I'll
leave it for you to figure out.
