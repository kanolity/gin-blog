package global

import (
	"github.com/sirupsen/logrus"
	"go_code/gin-vue-blog/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
