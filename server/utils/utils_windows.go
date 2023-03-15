//go:build windows

package utils

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

func Reg() {
	ex, err := os.Executable()
	PanicIfNeeded(err)

	capabilities, _, err := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\RemoteBrowser\Capabilities`, registry.ALL_ACCESS)
	PanicIfNeeded(err)
	defer capabilities.Close()
	err = capabilities.SetStringValue("ApplicationDescription", "open url on remote computer")
	PanicIfNeeded(err)
	err = capabilities.SetStringValue("ApplicationIcon", `%SystemRoot%\System32\SHELL32.dll,17`)
	PanicIfNeeded(err)
	err = capabilities.SetStringValue("ApplicationName", "RemoteBrowser")
	PanicIfNeeded(err)

	URLAssociations, _, err := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\RemoteBrowser\Capabilities\URLAssociations`, registry.ALL_ACCESS)
	PanicIfNeeded(err)
	defer URLAssociations.Close()
	err = URLAssociations.SetStringValue("ftp", "RemoteBrowserURL")
	PanicIfNeeded(err)
	err = URLAssociations.SetStringValue("http", "RemoteBrowserURL")
	PanicIfNeeded(err)
	err = URLAssociations.SetStringValue("https", "RemoteBrowserURL")
	PanicIfNeeded(err)

	registeredApplications, _, err := registry.CreateKey(registry.LOCAL_MACHINE, `SOFTWARE\RegisteredApplications`, registry.ALL_ACCESS)
	PanicIfNeeded(err)
	defer registeredApplications.Close()
	registeredApplications.SetStringValue("RemoteBrowser", `SOFTWARE\RemoteBrowser\Capabilities`)

	remoteBrowserURL, _, err := registry.CreateKey(registry.LOCAL_MACHINE, `Software\Classes\RemoteBrowserURL`, registry.ALL_ACCESS)
	PanicIfNeeded(err)
	defer remoteBrowserURL.Close()
	err = remoteBrowserURL.SetStringValue("", "RemoteBrowser Document")
	PanicIfNeeded(err)
	err = remoteBrowserURL.SetStringValue("FriendlyTypeName", "RemoteBrowser Document")
	PanicIfNeeded(err)

	shell, _, err := registry.CreateKey(registry.LOCAL_MACHINE, `Software\Classes\RemoteBrowserURL\shell`, registry.ALL_ACCESS)
	PanicIfNeeded(err)
	defer shell.Close()
	open, _, err := registry.CreateKey(registry.LOCAL_MACHINE, `Software\Classes\RemoteBrowserURL\shell\open`, registry.ALL_ACCESS)
	PanicIfNeeded(err)
	defer open.Close()
	command, _, err := registry.CreateKey(registry.LOCAL_MACHINE, `Software\Classes\RemoteBrowserURL\shell\open\command`, registry.ALL_ACCESS)
	PanicIfNeeded(err)
	defer command.Close()
	err = command.SetStringValue("", fmt.Sprintf(`"%s" -s "%s"`, ex, "%1"))
	PanicIfNeeded(err)
}
