package tag_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/res"
	"strconv"
)

func (api TagApi) TagListView(c *gin.Context) {
	// 获取分页和排序参数
	limit := c.DefaultQuery("limit", "10") // 每页数量，默认10
	page := c.DefaultQuery("page", "1")    // 当前页码，默认第1页
	sort := c.DefaultQuery("sort", "id")   // 排序字段，默认按ID排序

	// 转换分页参数
	var limitInt, pageInt int
	var err error
	if limitInt, err = strconv.Atoi(limit); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Invalid limit parameter"})
		return
	}
	if pageInt, err = strconv.Atoi(page); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Invalid page parameter"})
		return
	}
	offset := (pageInt - 1) * limitInt // 计算偏移量

	var tags []model.TagModel
	var totalCount int64

	// 查询标签及文章数量
	query := `
		SELECT 
			tag_model.*, 
			COUNT(article_tag_models.article_model_id) AS article_count,
			(SELECT COUNT(*) FROM tag_model) AS total_count
		FROM tag_model
		LEFT JOIN article_tag_models ON tag_model.id = article_tag_models.tag_model_id
		GROUP BY tag_model.id
		ORDER BY ? 
		LIMIT ? OFFSET ?
	`
	rows, err := global.DB.Raw(query, sort, limitInt, offset).Rows()
	if err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "Failed to fetch tags", "error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tag model.TagModel
		var articleCount int
		var rowCount int64
		if err := rows.Scan(&tag.ID, &tag.CreatedAt, &tag.UpdatedAt, &tag.Title, &articleCount, &rowCount); err != nil {
			c.JSON(500, gin.H{"code": 1, "msg": "Failed to scan row", "error": err.Error()})
			return
		}
		tag.ArticleCount = int64(articleCount)
		tags = append(tags, tag)
		totalCount = rowCount
	}

	if err := rows.Err(); err != nil {
		c.JSON(500, gin.H{"code": 1, "msg": "Failed to iterate rows", "error": err.Error()})
		return
	}

	res.OkWithList(tags, totalCount, c)
}
