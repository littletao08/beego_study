package entities
/**
 *文章
 */

type Article struct {
	Id int64
	UserId int64
	Title string
	Tag string
	CategoryId int32
	Content string
	CreatedAt string
	UpdatedAt string

}