// This file is generated, do not edit.
// Use 'fyne-icon-picker --regenerate' to generate this source file.

package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const IconPreviewSize = 96
var confirmCopyDialog dialog.Dialog

type iconPreviewLayout struct {
}

func (i iconPreviewLayout) Layout(objs []fyne.CanvasObject, parent fyne.Size) {
	objs[0].Resize(fyne.NewSize(IconPreviewSize, IconPreviewSize))
}
func (i iconPreviewLayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(IconPreviewSize, IconPreviewSize)
}

type tappableIcon struct {
	widget.Icon
	Text   string
	window fyne.Window
}

func newTappableIcon(resource fyne.Resource, text string, window fyne.Window) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(resource)
	icon.window = window
	icon.Text = text

	return icon
}

func (t *tappableIcon) Tapped(e *fyne.PointEvent) {
	previewIcon := widget.NewIcon(t.Resource)
	content := fyne.NewContainerWithLayout(
		iconPreviewLayout{},
		previewIcon,
	)
	copyOnConfirmCallback := func(accepted bool) {
		if accepted {
			t.window.Clipboard().SetContent("theme." + t.Text + "()")
		}
	}
	dialog.NewCustomConfirm(t.Text, "Copy", "Cancel", content, copyOnConfirmCallback, t.window)
}

type tappableLabel struct {
	widget.Label
	window fyne.Window
}

func newTappableLabel(text string, window fyne.Window) *tappableLabel {
	lbl := &tappableLabel{}
	lbl.ExtendBaseWidget(lbl)
	lbl.SetText(text)
	lbl.window = window

	return lbl
}

func (t *tappableLabel) Tapped(e *fyne.PointEvent) {
	t.window.Clipboard().SetContent("theme." + t.Text + "()")
	if confirmCopyDialog == nil {
		confirmCopyDialog = dialog.NewInformation("Copied", "Copied constant to clipboard", t.window)
	}
	confirmCopyDialog.Show()
	go func() {
		time.Sleep(time.Second)
		confirmCopyDialog.Hide()
	}()
}

func newIcon(resource fyne.Resource, name string, window fyne.Window) fyne.CanvasObject {
	icon := newTappableIcon(resource, name, window)
	label := newTappableLabel(name, window)
	label.window = window

	return fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
		icon,
		label,
	)
}

func getContents(win fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithColumns(6),
		newIcon(theme.FyneLogo(), "FyneLogo", win),
		newIcon(theme.CancelIcon(), "CancelIcon", win),
		newIcon(theme.ConfirmIcon(), "ConfirmIcon", win),
		newIcon(theme.DeleteIcon(), "DeleteIcon", win),
		newIcon(theme.SearchIcon(), "SearchIcon", win),
		newIcon(theme.SearchReplaceIcon(), "SearchReplaceIcon", win),
		newIcon(theme.MenuIcon(), "MenuIcon", win),
		newIcon(theme.MenuExpandIcon(), "MenuExpandIcon", win),
		newIcon(theme.CheckButtonIcon(), "CheckButtonIcon", win),
		newIcon(theme.CheckButtonCheckedIcon(), "CheckButtonCheckedIcon", win),
		newIcon(theme.RadioButtonIcon(), "RadioButtonIcon", win),
		newIcon(theme.RadioButtonCheckedIcon(), "RadioButtonCheckedIcon", win),
		newIcon(theme.ContentAddIcon(), "ContentAddIcon", win),
		newIcon(theme.ContentRemoveIcon(), "ContentRemoveIcon", win),
		newIcon(theme.ContentClearIcon(), "ContentClearIcon", win),
		newIcon(theme.ContentCutIcon(), "ContentCutIcon", win),
		newIcon(theme.ContentCopyIcon(), "ContentCopyIcon", win),
		newIcon(theme.ContentPasteIcon(), "ContentPasteIcon", win),
		newIcon(theme.ContentRedoIcon(), "ContentRedoIcon", win),
		newIcon(theme.ContentUndoIcon(), "ContentUndoIcon", win),
		newIcon(theme.DocumentIcon(), "DocumentIcon", win),
		newIcon(theme.DocumentCreateIcon(), "DocumentCreateIcon", win),
		newIcon(theme.DocumentPrintIcon(), "DocumentPrintIcon", win),
		newIcon(theme.DocumentSaveIcon(), "DocumentSaveIcon", win),
		newIcon(theme.InfoIcon(), "InfoIcon", win),
		newIcon(theme.QuestionIcon(), "QuestionIcon", win),
		newIcon(theme.WarningIcon(), "WarningIcon", win),
		newIcon(theme.FileIcon(), "FileIcon", win),
		newIcon(theme.FileApplicationIcon(), "FileApplicationIcon", win),
		newIcon(theme.FileAudioIcon(), "FileAudioIcon", win),
		newIcon(theme.FileImageIcon(), "FileImageIcon", win),
		newIcon(theme.FileTextIcon(), "FileTextIcon", win),
		newIcon(theme.FileVideoIcon(), "FileVideoIcon", win),
		newIcon(theme.FolderIcon(), "FolderIcon", win),
		newIcon(theme.FolderNewIcon(), "FolderNewIcon", win),
		newIcon(theme.FolderOpenIcon(), "FolderOpenIcon", win),
		newIcon(theme.HelpIcon(), "HelpIcon", win),
		newIcon(theme.HomeIcon(), "HomeIcon", win),
		newIcon(theme.SettingsIcon(), "SettingsIcon", win),
		newIcon(theme.MailAttachmentIcon(), "MailAttachmentIcon", win),
		newIcon(theme.MailComposeIcon(), "MailComposeIcon", win),
		newIcon(theme.MailForwardIcon(), "MailForwardIcon", win),
		newIcon(theme.MailReplyIcon(), "MailReplyIcon", win),
		newIcon(theme.MailReplyAllIcon(), "MailReplyAllIcon", win),
		newIcon(theme.MailSendIcon(), "MailSendIcon", win),
		newIcon(theme.MediaFastForwardIcon(), "MediaFastForwardIcon", win),
		newIcon(theme.MediaFastRewindIcon(), "MediaFastRewindIcon", win),
		newIcon(theme.MediaPauseIcon(), "MediaPauseIcon", win),
		newIcon(theme.MediaPlayIcon(), "MediaPlayIcon", win),
		newIcon(theme.MediaRecordIcon(), "MediaRecordIcon", win),
		newIcon(theme.MediaReplayIcon(), "MediaReplayIcon", win),
		newIcon(theme.MediaSkipNextIcon(), "MediaSkipNextIcon", win),
		newIcon(theme.MediaSkipPreviousIcon(), "MediaSkipPreviousIcon", win),
		newIcon(theme.MoveDownIcon(), "MoveDownIcon", win),
		newIcon(theme.MoveUpIcon(), "MoveUpIcon", win),
		newIcon(theme.NavigateBackIcon(), "NavigateBackIcon", win),
		newIcon(theme.NavigateNextIcon(), "NavigateNextIcon", win),
		newIcon(theme.MenuDropDownIcon(), "MenuDropDownIcon", win),
		newIcon(theme.MenuDropUpIcon(), "MenuDropUpIcon", win),
		newIcon(theme.ViewFullScreenIcon(), "ViewFullScreenIcon", win),
		newIcon(theme.ViewRestoreIcon(), "ViewRestoreIcon", win),
		newIcon(theme.ViewRefreshIcon(), "ViewRefreshIcon", win),
		newIcon(theme.ZoomFitIcon(), "ZoomFitIcon", win),
		newIcon(theme.ZoomInIcon(), "ZoomInIcon", win),
		newIcon(theme.ZoomOutIcon(), "ZoomOutIcon", win),
		newIcon(theme.VisibilityIcon(), "VisibilityIcon", win),
		newIcon(theme.VisibilityOffIcon(), "VisibilityOffIcon", win),
		newIcon(theme.VolumeDownIcon(), "VolumeDownIcon", win),
		newIcon(theme.VolumeMuteIcon(), "VolumeMuteIcon", win),
		newIcon(theme.VolumeUpIcon(), "VolumeUpIcon", win),
		newIcon(theme.ComputerIcon(), "ComputerIcon", win),
		newIcon(theme.DownloadIcon(), "DownloadIcon", win),
		newIcon(theme.StorageIcon(), "StorageIcon", win),
	)
}
