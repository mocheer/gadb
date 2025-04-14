package gadb

// Install
// adb install path_to_apk
// adb uninstall package_name
// func (d Device) Install(filename string) {

// }

// UninstallWithSystem
// 卸载系统应用
// adb shell pm uninstall --user 0 com.mumu.store
func (d Device) UninstallWithSystem(packageName string) (string, error) {
	return d.RunShellCommand("pm uninstall --user 0", packageName)
}

// 启动应用
// adb shell am start -n package_name/activity_name
func (d Device) Start(packageName string, activityName string) (string, error) {
	return d.RunShellCommand("am start -n", packageName+"/"+activityName)
}

// PackageInfo
// 查看包信息
// adb shell dumpsys package package_name
// 可以查看包的版本、Activity（用于启动）等信息
// android:name属性对应的值，即为启动Activity的完整类名。
func (d Device) PackageInfo(packageName string) (string, error) {
	return d.RunShellCommand("dumpsys package", packageName)
}
