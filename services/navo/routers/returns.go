package routers

import "time"

type AnnounceListResponse struct {
	Announces []struct {
		Body      string    `description:"内容"`   // 内容
		ClubID    string    `description:"社团ID"` // 社团ID
		CreatedAt time.Time `description:"创建时间"` // 创建时间
		ID        string    `description:"ID"`
		Kind      string    `description:"通过类型"` // 通过类型
		Name      string    `description:"通告名称"` // 通告名称
		UpdatedAt time.Time `description:"更新时间"` // 更新时间
		Receivers []struct {
			UserID string `description:"接受者的用户ID"` // 接受者的用户ID

		}
		Options struct {
			StartTime    time.Time `description:"开始时间"` // 开始时间
			EndTime      time.Time `description:"结束时间"` // 结束时间
			ReminderTime time.Time `description:"提醒时间"` // 提醒时间
		}
	}
	Total int32 `description:"总金额"`
}

type TemplatesResponse struct {
	Total     int32 `description:"总数"`
	Templates []struct {
		Title        string
		ClubID       string `description:"社团ID"`
		Description  string `description:"描述"`
		ReceiverInfo []struct {
			Avatar   string `description:"头像"`
			ID       string `description:"id"`
			Name     string `description:"名称"`
			Nickname string `description:"昵称"`
			RealName string `description:"真实姓名"`
		}
		Receivers []string `description:"接受者列表"`
		Config    struct {
			EndTime   string   `description:"结束时间"`
			StartTime string   `description:"开始时间"`
			Kind      string   `description:"种类`
			Weeks     []string `description:"周列表, 0:日,1:,2,3,4,5,6:六"`
		} `description:"配置信息`
		Creator struct {
			Avatar   string
			ID       string
			Name     string
			Nickname string
			RealName string
		}
		CustomFields []struct {
			Key     string `description:"健值"`
			Kind    string `description:"种类"`
			Label   string `description:"标签"`
			Options struct {
				config []struct {
					name  string `description:"选项名称"`
					value string `description:"选项值"`
				} `description:"选项配置"`
			} `description:"字段配置"`
		}
	}
}
