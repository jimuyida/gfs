#
golang fastdfs client


STORAGE_PROTO_CMD_APPEND_FILE  直接追加文件末尾
STORAGE_PROTO_CMD_MODIFY_FILE  有offset追加

```golang
client,e := gfs.NewFdfsClient([]string{"192.168.0.147:22122"})
```

```golang
r,e := client.UploadByFilename(filename)
```

```golang
client.DeleteFile(r.FileId)
```

```golang
    r,e := client.UploadAppenderByFilename(filename)
	log.Println(r,e)
	e = client.ModifyByFilename(filename,r.FileId,15717)
	log.Println(r,e)
```

```golang
    f,e := client.QueryFile(r.FileId)
	log.Println(f,e)
```

```golang
    //group1/M00/00/00/wKgAk1oVPICASFqAAAA9Zbiunec638.jpg
    r1,e := client.DownloadToFile("1.jpg",r.FileId,0,0)
    log.Println(r1,e)
```