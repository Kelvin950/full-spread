package domain 


type ApiError struct{

	Code int 
	ErrVal error
}



func(a  ApiError) Error()string {


return a.ErrVal.Error()

}
