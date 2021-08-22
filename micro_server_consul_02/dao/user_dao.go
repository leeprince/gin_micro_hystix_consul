package dao

import (
	"gorm.io/gorm"
	"micro_server/models"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/13 下午11:53
 * @Desc:
 */

type UserDao struct {
	DB     *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{
		DB: models.DB(),
	}
}

type UsersOpt struct {
	models.Users
	Select []string
	Where  string
}

func (d *UserDao) GetUserByUserId(userId int64) (*models.Users, error) {
	opt := UsersOpt{}
	opt.Id = userId
	opt.Select = []string{"username"}
	user := new(models.Users)
	
	// err := d.DB.Select(d.Select).Where(d.Users).First(&user).Error
	//.Debug()开启debug模式在控制台输出sql
	// err := d.DB.Debug().Select(d.Select).Where(d.Users).Find(&user).Error
	//.Unscoped()忽略：`deleted_at` IS NULL
	err := d.DB.Debug().Select(opt.Select).Where(opt.Users).Unscoped().Find(&user).Error
	// err := d.DB.Debug().Where(opt.Users).Unscoped().Find(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
