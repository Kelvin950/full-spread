package server

type CreateMultiPartUploadReq struct {
	Key        string `json:"key" binding:"required"`
	BucketName string `json:"bucket_name" binding:"required"`
}

type UplaodMultiPartReq struct {
	Bucket     string `json:"bucket" binding:"required"`
	UploadId   string `json:"upload_id" binding:"required"`
	Key        string `json:"key" binding:"required"`
	PartNumber []int  `json:"part_number" binding:"required"`
}
