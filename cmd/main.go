package main

import (
	"../../gfs"
	"log"
)

type AuthSuccess struct {
	/* variables */
	AccessToken string  `json:"access_token"`
	TokenType string `json:"token_type"`
	Scope string `json:"scope"`
	ExpiresIn  int64 `json:"expires_in"`
}

func main() {
	client,e := gfs.NewFdfsClient([]string{"192.168.0.147:22122"})

	log.Println(client,e)
	filename := "C:\\Users\\WXH\\Desktop\\lgo.jpg"
	//filename := "C:\\Users\\WXH\\Downloads\\nox_setup_v6.0.0.0_full.exe"
	r,e := client.UploadByFilename(filename)
	log.Println(r,e)
	/*
	组名：group1
	磁盘：M00
	目录：00/00
	文件名称：wKiAg1lE9WqAWu_ZAAFaL_xdW_s943.jpg
	*/

	defer client.DeleteFile(r.FileId)

	r,e = client.UploadAppenderByFilename(filename)
	log.Println(r,e)
	e = client.ModifyByFilename(filename,r.FileId,15717)
	log.Println(r,e)

	f,e := client.QueryFile(r.FileId)
	log.Println(f,e)

	r1,e := client.DownloadToFile("1.jpg",r.FileId,0,0)
	log.Println(r1,e)

}
