package models

// IntegralSourceType 定义积分日志来源的枚举类型
type IntegralSourceType int

const (
	// SourceTypePost 发帖
	SourceTypePost IntegralSourceType = iota + 1 // 发帖
	// SourceTypeComment 评论
	SourceTypeComment // 评论
	// SourceTypeReceivedComment 被评论
	SourceTypeReceivedComment // 被评论
	// SourceTypeSignIn 签到
	SourceTypeSignIn // 签到
	// SourceTypeRecharge 充值
	SourceTypeRecharge // 充值
	// SourceTypePurchaseAttachment 购买附件
	SourceTypePurchaseAttachment // 购买附件
	// SourceTypeReward 打赏
	SourceTypeReward // 打赏
	// SourceTypeBounty 悬赏
	SourceTypeBounty // 悬赏
)
const (
	// IntegralTypeAdd 增加积分
	IntegralTypeAdd = 1
	// IntegralTypeReduce 减少积分
	IntegralTypeReduce = 2
)

// IntegralSourceTypeToString 将IntegralSourceType转换为字符串
func IntegralSourceTypeToString(s IntegralSourceType) string {
	switch s {
	case SourceTypePost:
		return "发帖"
	case SourceTypeComment:
		return "评论"
	case SourceTypeReceivedComment:
		return "被评论"
	case SourceTypeSignIn:
		return "签到"
	case SourceTypeRecharge:
		return "充值"
	case SourceTypePurchaseAttachment:
		return "购买附件"
	case SourceTypeReward:
		return "打赏"
	case SourceTypeBounty:
		return "悬赏"
	default:
		return "未知"
	}
}

// 系统配置
const (
	SiteName        = "siteName"        // 网站名称
	SiteHost        = "siteHost"        // 网站域名
	SiteDescription = "siteDescription" //	网站描述
	SiteKeywords    = "siteKeywords"    //	网站关键字
	SiteNotice      = "siteNotice"      //	网站公告
)

// GrouplistKey 分组Rediskey
const (
	GroupListKey = "groupList" // 分组列表
)
