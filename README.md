# fyne-icon-picker
An interactive icon picker. The options are generated from those included in the theme package of fyne.

The motivation was two-fold:
1. I wanted an interactive experience that allowed me to preview the icons I was including in my fyne UI code.
1. I didn't want such a tool to be hard to use or maintain.

So with a bit of regex the fyne/theme/icons.go file is parsed for icon functions that are then injected into
the template generating the `generated.go` source file... simple.

## Regeneration

This assumes that it can find your fyne source download at `~/go/fyne.io/fyne/theme/icons.go`. If that's not
the case, this will not be able to regenerate from source. I could probably look for an environment defined
path, or just pull the latest source file from Github, but this works for now.

```bash
# Regenerate the generated.go source file from your local fyne download.
fyne-icon-picker --regenerate
```
