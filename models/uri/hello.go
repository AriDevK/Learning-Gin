package uri

type HelloUri struct {
	Name  string `uri:"name" binding:"required"`
	Times int    `uri:"times" binding:"required"`
}
