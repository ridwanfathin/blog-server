package post

import (
	"github.com/ridwanfathin/blog-server/api/tag"

	"github.com/jinzhu/gorm"
)

func AddTag(db *gorm.DB, post *Post, tag *tag.Tag) error {
	res := db.Model(post).Association(AssociationTags).Append(tag)
	return res.Error
}
