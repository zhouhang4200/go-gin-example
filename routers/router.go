package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zhouhang4200/go-gin-example/docs"
	"github.com/zhouhang4200/go-gin-example/middleware/jwt"
	"github.com/zhouhang4200/go-gin-example/pkg/export"
	"github.com/zhouhang4200/go-gin-example/pkg/setting"
	"github.com/zhouhang4200/go-gin-example/pkg/upload"
	"github.com/zhouhang4200/go-gin-example/routers/api"
	"github.com/zhouhang4200/go-gin-example/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath())) // 查看图片
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//apiv1.GET("/article", v1.GetArticle)
		//apiv1.GET("/articles/:id", v1.GetArticles)
		//apiv1.POST("/articles", v1.AddArticle)
		//apiv1.PUT("/articles/:id", v1.EditArticle)
		//apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
