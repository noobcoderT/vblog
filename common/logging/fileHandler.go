package logging

import (
    "os"
    "path"
    "log"
    "vblog/common/setting"
)

var (
    logFilePath = setting.LogFilePath
    logFileName = setting.LogFileName
    fileName = path.Join(logFilePath, logFileName)
    workDir, _ = os.Getwd()
    fullLogFilePath = workDir + "/" + fileName
)

func openLogFile() (*os.File, error) {
    _, err := os.Stat(logFilePath)
    switch {
        case os.IsNotExist(err):
            mkDir()
        case os.IsPermission(err):
            log.Fatalf("Permission :%v", err)
    }

    handle, err := os.OpenFile(fileName, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Fail to OpenFile :%v", err)
    }

    return handle, err
}

func mkDir() {
    dir, _ := os.Getwd()
    err := os.MkdirAll(dir + "/" + logFilePath, os.ModePerm)
    if err != nil {
        panic(err)
    }
}
