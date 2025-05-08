package domain 


type CreateMultiPartUpload struct{
Key *string  `json:"key"`
BucketName *string  `json:"bucket_name"`

}


type  UplaodMultiPart struct {
Bucket  *string `json:"bucket"`
UploadId  *string  `json:"upload_id"` 
Key *string `json:"key"`
PartNumber *int32 `json:"part_number"`
}


type  CompleteMultiPart struct{
Bucket  *string `json:"bucket"`
UploadId  *string  `json:"upload_id"` 
Key *string `json:"key"`
*MultipartUpload
}

type MultipartUpload struct{

	Part  []Parts
}

type Parts struct{
	Etag *string `json:"etag"` 
	PartNumber *int32 `json:"part_number"`
}