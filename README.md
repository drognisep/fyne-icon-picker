# fyne-icon-picker
An interactive icon picker. The options are generated from those included in the theme package of fyne.

The motivation was two-fold:
1. I wanted an interactive experience that allowed me to preview the icons I was including in my fyne UI code.
1. I didn't want such a tool to be hard to use or maintain.

So with a bit of regex the fyne/theme/icons.go file is parsed for icon functions that are then injected into
the template generating the `generated.go` source file... simple.

![Initial State](https://github.com/drognisep/fyne-icon-picker/blob/master/docs/images/initial.png?raw=true)

## Preview Icons
Click on any icon to show a preview dialog, then choose whether to copy the fyne function name to reference that
icon, or hit Cancel to close it.

![Icon Preview](https://github.com/drognisep/fyne-icon-picker/blob/master/docs/images/icon_clicked.png?raw=true)

## Copy Constant Directly
Clicking on the text identifier to the right of an icon will copy the fyne function name directly, and a
dialog will confirm that interaction.

![Copied!](https://github.com/drognisep/fyne-icon-picker/blob/master/docs/images/copied.png?raw=true)

## Regeneration

This assumes that it can find your fyne source download at `~/go/fyne.io/fyne/theme/icons.go`. If that's not
the case, this will not be able to regenerate from source. I could probably look for an environment defined
path, or just pull the latest source file from Github, but this works for now.

```bash
# Regenerate the generated.go source file from your local fyne download.
fyne-icon-picker --regenerate
```
