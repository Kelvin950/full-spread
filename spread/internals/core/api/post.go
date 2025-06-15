package api

import "github.com/kelvin950/spread/internals/core/domain"


func(a Api)GetCreatorPosts(userId ,page , pageSize int)([]domain.Post , error){

  creator , err:= 	a.Db.GetCreator(domain.Creator{
		UserID: uint(userId),
	})

	if err!=nil{
		return nil , err 
	}

	return a.Db.GetCreatorPosts(creator.ID , page, pageSize)
}


func(a Api)GetCreatorPost(userID , postid int)(domain.Post , error){

	creator ,err:= a.Db.GetCreator(domain.Creator{
		UserID: uint(userID),
	})

	if err!=nil{
		return domain.Post{}, err
	}

	return a.Db.GetCreatorPost(creator.ID ,uint(postid)) 

}