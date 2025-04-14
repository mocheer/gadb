package gadb

import (
	"path"
	"path/filepath"
)

// 支持修改默认的共享目录
var SharedMuMu = `C:\Users\Administrator\Documents\MuMu共享文件夹`

// screencap:
// usage: screencap [-hp] [-d display-id] [FILENAME]
//    -h: this message
//    -p: save the file as a png.
//    -d: specify the physical display ID to capture (default: 4619827820427265280)
//        see "dumpsys SurfaceFlinger --display-id" for valid display IDs.
// If FILENAME ends with .png it will be saved as a png.
// If FILENAME is not given, the results will be printed to stdout.

// Screencap
// adb screencap -p --crop=w:h:x:y
// 一般要带上根目录：sdcard
// d.Screencap("/sdcard/a.png")
func (d Device) Screencap(filename string) (string, error) {
	return d.RunShellCommand("screencap", filename)
}

// ScreencapSaveSD
func (d Device) ScreencapSaveSD(filename string) (string, error) {
	return d.RunShellCommand("screencap", path.Join("/sdcard", filename))
}

// ScreencapSaveMuMuShared
// 直接截图到共享目录下的Screenshots目录，方便读取
func (d Device) ScreencapSaveMuMuShared(filename string) (string, error) {
	return d.RunShellCommand("screencap", path.Join("/sdcard/$MuMu12Shared/Screenshots", filename))
}

// GetWindoweMuMuSharedPath
func (d Device) GetWindoweMuMuSharedScreenshotsPath(filename string) string {
	return filepath.Join(SharedMuMu, "Screenshots")
}

// 一下代码无法运行
// // ExecOutScreencap
// // adb exec-out screencap -p
// // 相对于当前目录
// func (d Device) ExecOutScreencap(filename string) (string, error) {
// 	return d.RunCommand("exec:exec-out:", "screencap", filename)
// }

// func (d Device) RunCommand(cmd string, args ...string) (string, error) {
// 	raw, err := d.RunCommandWithBytes(cmd, args...)
// 	return string(raw), err
// }

// func (d Device) RunCommandWithBytes(cmd string, args ...string) ([]byte, error) {
// 	if len(args) > 0 {
// 		cmd = fmt.Sprintf("%s %s", cmd, strings.Join(args, " "))
// 	}
// 	if strings.TrimSpace(cmd) == "" {
// 		return nil, errors.New("adb: command cannot be empty")
// 	}
// 	log.Println(cmd)
// 	raw, err := d.executeCommand(cmd)
// 	return raw, err
// }
