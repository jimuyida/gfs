# gfs
golang fastdfs client

修改自：https://github.com/sanxia/gfs

STORAGE_PROTO_CMD_APPEND_FILE  直接追加文件末尾
STORAGE_PROTO_CMD_MODIFY_FILE  有offset追加

### 连接
```golang
client,e := gfs.NewFdfsClient([]string{"192.168.0.147:22122"})
```

### 上传文件
```golang
r,e := client.UploadByFilename(filename)
```

### 删除文件
```golang
client.DeleteFile(r.FileId)
```

### 追加文件
```golang
r,e := client.UploadAppenderByFilename(filename)
log.Println(r,e)
e = client.ModifyByFilename(filename,r.FileId,15717)
log.Println(r,e)
```

### 获取文件信息
```golang
f,e := client.QueryFile(r.FileId)
log.Println(f,e)
```
### 下载文件
```golang
//group1/M00/00/00/wKgAk1oVPICASFqAAAA9Zbiunec638.jpg
r1,e := client.DownloadToFile("1.jpg",r.FileId,0,0)
log.Println(r1,e)
```
