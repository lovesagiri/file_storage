package conf

//文件状态
//1 不存在   2 等待上传  3 正在上传  4 存在  5 过期
const (
	FileNotExist = iota + 1
	FileWaitingForUpload
	FileUploading
	FileAlreadyExists
	FileExpired
)