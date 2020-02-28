package router

import (
    //"net/http"
    "github.com/gin-gonic/gin"
    "vblog/router/api"
    "vblog/router/api/v1"
    "vblog/common/setting"
    "vblog/middleware/jwt"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    gin.SetMode(setting.RunMode)
    r.GET("/auth", api.GetAuth)
    apiv1 := r.Group("/api/v1")
    apiv1.Use(jwt.JWT())
    {
        //获取标签列表
        apiv1.GET("/tag", v1.GetTags)
        //新建标签
        apiv1.POST("/tag", v1.AddTag)
        //更新指定标签
        apiv1.PUT("/tag/:id", v1.EditTag)
        //删除指定标签
        apiv1.DELETE("/tag/:id", v1.DeleteTag)
        //获取文章列表
        apiv1.GET("/article", v1.GetArticles)
        //获取指定文章
        apiv1.GET("/article/:id", v1.GetArticle)
        //新建文章
        apiv1.POST("/article", v1.AddArticle)
        //更新指定文章
        apiv1.PUT("/article/:id", v1.EditArticle)
        //删除指定文章
        apiv1.DELETE("/article/:id", v1.DeleteArticle)
    }

    return r
}
