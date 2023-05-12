package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			return
		}
		//判断文件是否存在
		bytes := make([]byte, header.Size)
		//将文件信息写入byte
		_, err = file.Read(bytes)
		if err != nil {
			return
		}
		//生成16进制数 转字符串  根据文件信息生成唯一hash , 相同的文件生成的hash是相同的
		hash := fmt.Sprintf("%x", md5.Sum(bytes))
		rp := new(models.RepositoryPool)
		//是否存在这个hash
		get, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if get {
			httpx.OkJson(w, &types.FileUploadResponse{Identity: rp.Identity})
			return
		}

		//文件不存在  将客户端上传的文件  往cos存储文件
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}
		//复制到req FileUploadRequest   往logic传
		req.Name = header.Filename
		req.Ext = path.Ext(header.Filename)
		req.Size = header.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
