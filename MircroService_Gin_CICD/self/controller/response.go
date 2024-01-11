package controller

import (
	"MicroService/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	// 取文章列表(应该是从数据库中取的，这里简化了就在内存中取了)
	articles := GetAllArticles()
	// code表示HTTP状态码，name表示要渲染的模板文件名，obj表示要传递给模板的数据对象
	// 使用context.HTML的方法来渲染
	c.HTML(
		http.StatusOK,
		"index.html",
		// 将页眉模板中title设置为HomePage
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

// 相比于之前改进使用render函数，自动确定信息的渲染形式
func Home_new(c *gin.Context){
	articles := GetAllArticles()
	data := gin.H{
		"title":   "Home Page",
		"payload": articles,
	}
	Render(c,data,"index.html")
}

func GetAllArticles() []models.Article {
	return models.ArticleList
}

func GetArticle(c *gin.Context) {
	// c.Param获得的关键字段都是string，需要转为int
	if articleID, err := strconv.Atoi(c.Param("article_id"));err ==nil{
		if article,err := GetArticleByID(articleID);err==nil{
			// 将信息传入HTML中进行渲染
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":article.Title,
					"payload":article,
				},
			)
		}else{
			c.AbortWithError(http.StatusNotFound, err)
		}
	}else{
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func GetArticleByID(id int) (*models.Article, error) {
	for _, article := range models.ArticleList {
		if article.ID == id {
			return &article, nil
		}
	}
	return nil, errors.New("Article not found")
}
