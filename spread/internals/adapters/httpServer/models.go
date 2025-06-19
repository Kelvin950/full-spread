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

type CompleteMultiPartReq struct {
	Bucket   string  `json:"bucket"`
	UploadId string  `json:"upload_id"`
	Key      string  `json:"key"`
	Parts    []Parts `json:"parts"`
}

type Parts struct {
	Etag       *string `json:"etag"`
	PartNumber *int32  `json:"part_number"`
}

type CreateUserReq struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	DateOfBirth string `json:"dob"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginGoogleUserReq struct {
	Credential string `json:"credential"`
}

type CreateCreatorReq struct {
	Name        string `json:"name"`
	HeaderUrl   string `json:"header_url"`
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateCreatorReq struct {
	Id          uint   `json:"creator_id"`
	Name        string `json:"name"`
	HeaderUrl   string `json:"header_url"`
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
}

type CreateMemberReq struct {
	UserID    uint `json:"user_id"`
	CreatorId uint `json:"creator_id"`
}

type CreateSubReq struct {
	MembershipID uint `json:"membership_id"`
}
type CreatePostReq struct {
	Description string `json:"description"`
	Type        string `json:"type"`
	Topics      []uint `json:"topics"`
	Content    []CreateContentReq `json:"content"`
}

type CreateContentReq struct {
	MimeType        string  `json:"mime_type"`
	LocationUrl     string  `json:"location_url"`
}