package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"goblog-backend/db"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {
	// 自定义规则 is_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// is_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// is_exists:users,email 检查数据库表里是否存在同一条信息
	// is_exists:users,email,2 排除用户掉 id 为 2 的用户
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

	// max_cn:8 中文长度设定不超过 8
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}
		return nil
	})

	// min_cn:2 中文长度设定不小于 2
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})

	// 自定义规则 not_exists，确保数据库存在某条数据
	// 一个使用场景是创建话题时需要附带 category_id 分类 ID 为参数，此时需要保证
	// category_id 的值在数据库中存在，即可使用：
	// not_exists:categories,id
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 categories
		tableName := rng[0]
		// 第二个参数，字段名称，如 id
		dbFiled := rng[1]

		// 用户请求过来的数据
		requestValue := value.(string)

		// 查询数据库
		var count int64
		db.Db.Table(tableName).Where(dbFiled+" = ?", requestValue).Count(&count)
		// 验证不通过，数据不存在
		if count == 0 {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 不存在", requestValue)
		}
		return nil
	})
}
