package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goblog-backend/db"
	"strings"
)

func init() {
	// 自定义规则 is_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,2 排除用户掉 id 为 2 的用户
	govalidator.AddCustomRule("is_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "is_exists:"), ",")

		tableName := rng[0]
		dbField := rng[1]

		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		requestValue := value.(string)

		query := db.Db.Table(tableName).Where(dbField+" =?", requestValue)

		if len(exceptID) > 0 {
			query = query.Where("id!=?", exceptID)
		}

		var count int64
		query.Count(&count)

		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", requestValue)
		}

		return nil
	})
}
