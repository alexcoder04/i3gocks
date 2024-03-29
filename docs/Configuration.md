
# General

Configuration is done in [yaml](https://yaml.org/), a human- and machine-friendly markup language.

The config file is located at `$XDG_CONFIG_HOME/i3gocks/config.yml`

If the config does not exist or cannot be loaded, the minimal default hardcoded config is used.

# Config structure

```yml
Options:
  ...
Modules:
 - ...
 - ...
Colors:
 - ...
 - ...
```

# Options

```yml
Options:
  PowerlineTheme: true
  PowerlineSeparator: "\uE0BE"
```

If `PowerlineTheme` is set to true, i3gocks will draw a custom separator between
the blocks. This separator is defined in `PowerlineSeparator` and is supposed to
be a Nerd Font / Powerline glyph (i3gocks will set its foreground color to
module's background color, and the background to previous module's background).

# Modules

## Example with all fields (only Name and Command are mandatory):

```yml
 - Name: time
   ForegroundColor: "*green"   # default: "*white"
   BackgroundColor: "#111111"  # default: "*black"
   Pre: "⏰ "                  # default: ""
   Post: " "                   # default: ""
   Command: "date"
   Args: ["+%d.%m.%Y - %R:%S"] # default: []
   Interval: 1                 # default: 1
   Markup: "none"              # default: "none"
   Signal: 2                   # default: 0
```

## Name

Key that is used by i3bar/swaybar and i3gocks to identify the blocks.

Possible values: any string

## ForegroundColor

Text color of the block.

Possible values: any color in hex notation (`#123456`) / color reference (`*red`). If color reference is used, the repective color is loaded from the Colors section.

## BackgroundColor

Background color of the block.

Possible values: see ForegroundColor

## Pre

String added before the block text.

Possible values: any string.

## Post

String added after the block text.

Possible values: any string.

## Command

Command that is executed to get the block text. The output lines are read similar to i3blocks:

```text
Block text (full_text)
-ignored-
ForegroundColor
BackgroundColor
```

If the command starts with `*`, it is interpreted as a name of a built-in module. Those are:

 - `*time`: the current date and time. Args: `["format"]`.
   Format is in the Go time format fashion (`02.01.2006 (Mon) 15:04:05`).
 - `*echo`: print some text. Args: `["text", "text", ...]`

Possible values: any executable on your system / any builtin.

## Args

Command-line arguments that are passed to the command.

Possible values: any array of strings.

## Interval

How often (in seconds) the command output is re-loaded.

Possible values: any positive integer.

## Markup

Whether i3bar should format the output using markup.

Possible values: `none` / `pango`

## Signal

Number of signal (`SIGRTMIN+n`) that can be sent to i3gocks to reload the block.

Possible values: `1`-`15`

# Colors

The colors array is optional, colors are loaded from these locations before (later locations override previous ones):

1. Hardcoded defaults (gruvbox dark theme).
2. Environment. Any environmental variable that starts with `COLOR_` is considered a color value. Example: `$COLOR_RED` is loaded as `RED`
3. Your config

