package larkt

import (
	"context"
	"fmt"
	"net/http"
	"time"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/pkg/errors"
)

type DBTicket struct {
	ID            int64     `gorm:"column:id;primaryKey;comment:'主键,自增长'"` // 主键
	TenantID      int64     `gorm:"column:tenant_id;comment:'租户ID'"`
	CanaryPercent int32     `gorm:"column:canary_percent;comment:'灰度发布比例'"`
	ModuleName    string    `gorm:"column:module_name;type:varchar(32);comment:'模块名称'"`
	Creator       string    `gorm:"column:creator;type:varchar(64);comment:'创建者'"`
	Reviewer      string    `gorm:"column:reviewer;type:varchar(64);comment:'审批者'"`
	Before        string    `gorm:"column:before;type:text;comment:'修改前的rule快照'"`
	After         string    `gorm:"column:after;type:text;comment:'修改后的rule内容'"`
	TicketStatus  string    `gorm:"column:ticket_status;type:varchar(16);comment:'工单状态'"`
	ReviewStatus  string    `gorm:"column:review_status;type:varchar(16);comment:'审批状态'"`
	WhiteList     string    `gorm:"column:white_list;type:varchar(2048);comment:'设备ID白名单'"`
	Description   string    `gorm:"column:description;type:varchar(256);comment:'工单的描述性文字'"`
	CreatedAt     time.Time `gorm:"column:created_at;comment:'创建时间,自动填充'"`
	ReviewedAt    time.Time `gorm:"column:reviewed_at;comment:'审批时间,审批时更新'"`
}
type KerLarkRobot struct {
	client *lark.Client
}

var LarkRobot *KerLarkRobot

func InitLarkRobot() {
	LarkRobot = &KerLarkRobot{
		client: lark.NewClient("cli_a5ad2707dccd500c", "WR0UnPtakkuRmQh6HbKVsgGlwVw8Mtpr",
			lark.WithLogLevel(larkcore.LogLevelDebug),
			lark.WithReqTimeout(3*time.Second),
			lark.WithEnableTokenCache(true),
			// lark.WithHelpdeskCredential("id", "token"),
			lark.WithHttpClient(http.DefaultClient)),
	}
}

// SendLarkCardMessageToReceiver 给接受者发送消息(message 来自LarkNotificationMessageCardXXX函数)
func (robot *KerLarkRobot) SendLarkCardMessageToReceiver(receiver, message string) {
	msgBody := larkim.NewCreateMessageReqBodyBuilder().
		ReceiveId(receiver + "@bytedance.com").
		MsgType(larkim.MsgTypeInteractive).
		Content(message).
		Build()

	req := larkim.NewCreateMessageReqBuilder().ReceiveIdType(larkim.ReceiveIdTypeEmail).Body(msgBody).Build()
	if resp, err := robot.client.Im.Message.Create(context.Background(), req); err != nil || !resp.Success() {
		return
	}
}

// LarkNotificationMessageCard4RequestReview Request Review消息卡，用于给reviewer发请求通知
func LarkNotificationMessageCard4RequestReview(ticket *DBTicket) (string, error) {
	if ticket == nil {
		return "", errors.New("input is nil")
	}

	larkCardLayout := larkcard.NewMessageCard().
		Config(NewCardConfig()).                // 卡片整体配置
		Header(NewCardHeaderTitle()).           // 卡片的title
		Elements([]larkcard.MessageCardElement{ // 卡片正文布局
			// 第1行的信息
			larkcard.NewMessageCardDiv().Fields([]*larkcard.MessageCardField{
				NewCardField4Creator(ticket.Creator, true),
				NewCardField4CreationTime(ticket.CreatedAt, true),
			}).Build(),

			// 第2行的信息
			larkcard.NewMessageCardDiv().Fields([]*larkcard.MessageCardField{
				NewCardField4Reviewer(ticket.Reviewer, true),
				NewCardField4ReviewStatus(ticket.ReviewStatus, true),
			}).Build(),

			// 第3行的信息
			larkcard.NewMessageCardDiv().Fields([]*larkcard.MessageCardField{
				NewCardField4Description(ticket.Description, false),
			}).Build(),

			NewCardDividingLine(), // 第四行 分割线

			// 第五行 交互按钮布局
			larkcard.NewMessageCardAction().Actions([]larkcard.MessageCardActionElement{
				NewCardButton4Reviewer(ticket.ID),
			}).Build(),
		}).Build()

	return larkCardLayout.JSON()
}

// LarkNotificationMessageCard4Reviewed  Reviewed消息卡，用于通知创建者审批结果
func LarkNotificationMessageCard4Reviewed(ticket *DBTicket) (string, error) {
	if ticket == nil {
		return "", errors.New("input is nil")
	}

	larkCardLayout := larkcard.NewMessageCard().
		Config(NewCardConfig()).                // 卡片整体配置
		Header(NewCardHeaderTitle()).           // 卡片的title
		Elements([]larkcard.MessageCardElement{ // 卡片正文布局
			// 第1行的信息
			larkcard.NewMessageCardDiv().Fields([]*larkcard.MessageCardField{
				NewCardField4Creator(ticket.Creator, true),
				NewCardField4CreationTime(ticket.CreatedAt, true),
			}).Build(),

			// 第2行的信息
			larkcard.NewMessageCardDiv().Fields([]*larkcard.MessageCardField{
				NewCardField4Reviewer(ticket.Reviewer, true),
				NewCardField4ReviewStatus(ticket.ReviewStatus, true),
			}).Build(),

			// 第3行的信息
			larkcard.NewMessageCardDiv().Fields([]*larkcard.MessageCardField{
				NewCardField4ReviewedTime(ticket.ReviewedAt, true),
				NewCardField4ReviewStatus(ticket.ReviewStatus, true),
			}).Build(),

			// 第4行的信息
			larkcard.NewMessageCardDiv().Fields([]*larkcard.MessageCardField{
				NewCardField4Description(ticket.Description, false),
			}).Build(),

			NewCardDividingLine(), // 第五行 分割线

			// 第六行 交互按钮布局
			larkcard.NewMessageCardAction().Actions([]larkcard.MessageCardActionElement{
				NewCardButton4Creator(ticket.ID),
			}).Build(),
		}).Build()

	return larkCardLayout.JSON()
}

// NewCardDividingLine card中的分割线
func NewCardDividingLine() *larkcard.MessageCardHr {
	return larkcard.NewMessageCardHr().Build()
}

// NewMessageCardFieldWithMarkdown 创建一个md格式的card field消息，isShort：一行对列则为true，否则为false
func NewMessageCardFieldWithMarkdown(content string, isShort bool) *larkcard.MessageCardField {
	return larkcard.NewMessageCardField().
		IsShort(isShort).
		Text(larkcard.NewMessageCardLarkMd().Content(content).Build()).
		Build()
}

// NewMessageCardEmbedButton 创建一个button并携带URL
func NewMessageCardEmbedButton(text string, url string) *larkcard.MessageCardEmbedButton {
	return larkcard.NewMessageCardEmbedButton().
		Type(larkcard.MessageCardButtonTypePrimary).Url(url).
		Text(larkcard.NewMessageCardPlainText().Content(text).Build()).
		Build()
}

func NewMessageCardHeaderTitle(color, content string) *larkcard.MessageCardHeader {
	return larkcard.NewMessageCardHeader().
		Template(color).
		Title(larkcard.NewMessageCardPlainText().Content(content).Build()).
		Build()
}

// NewCardConfig card configuration
func NewCardConfig() *larkcard.MessageCardConfig {
	return larkcard.NewMessageCardConfig().EnableForward(true).Build()
}

// NewCardHeaderTitle new card title
func NewCardHeaderTitle() *larkcard.MessageCardHeader {
	return NewMessageCardHeaderTitle(larkcard.TemplateRed, "VOD Strategy Platform")
}

// NewCardField4Creator creator info
func NewCardField4Creator(creator string, isShort bool) *larkcard.MessageCardField {
	return NewMessageCardFieldWithMarkdown(fmt.Sprintf("Creator:\n%s", creator), isShort)
}

// NewCardField4Reviewer reviewer info
func NewCardField4Reviewer(reviewer string, isShort bool) *larkcard.MessageCardField {
	return NewMessageCardFieldWithMarkdown(fmt.Sprintf("Reviewer:\n%s", reviewer), isShort)
}

// NewCardField4CreationTime creation time
func NewCardField4CreationTime(createdAt time.Time, isShort bool) *larkcard.MessageCardField {
	// TODO 变换时间格式
	return NewMessageCardFieldWithMarkdown(fmt.Sprintf("Creation Time:\n%s", createdAt), isShort)
}

// NewCardField4ReviewedTime reviewed timestamp
func NewCardField4ReviewedTime(reviewedAt time.Time, isShort bool) *larkcard.MessageCardField {
	// TODO 变换时间格式
	return NewMessageCardFieldWithMarkdown(fmt.Sprintf("Reviewed Time:\n%s", reviewedAt), isShort)
}

// NewCardField4TenantName tenant name info
func NewCardField4TenantName(tenantName string, isShort bool) *larkcard.MessageCardField {
	return NewMessageCardFieldWithMarkdown(fmt.Sprintf("Tenant Name:\n%s", tenantName), isShort)
}

// NewCardField4ReviewStatus review status info
func NewCardField4ReviewStatus(reviewStatus string, isShort bool) *larkcard.MessageCardField {
	return NewMessageCardFieldWithMarkdown(fmt.Sprintf("Review Status:\n%s", reviewStatus), isShort)
}

// NewCardField4Description description of ticket
func NewCardField4Description(description string, isShort bool) *larkcard.MessageCardField {
	return NewMessageCardFieldWithMarkdown(fmt.Sprintf("Description:\n%s", description), isShort)
}

// NewCardButton4Reviewer view and review button
func NewCardButton4Reviewer(ticketID int64) *larkcard.MessageCardEmbedButton {
	url := fmt.Sprintf("https://vod-strategy.sre-boei18n.bytedance.net/tickets/details?ticket_id=%d", ticketID)
	return NewMessageCardEmbedButton("View and Review", url)
}

// NewCardButton4Creator view button
func NewCardButton4Creator(ticketID int64) *larkcard.MessageCardEmbedButton {
	url := fmt.Sprintf("https://vod-strategy.sre-boei18n.bytedance.net/tickets/details?ticket_id=%d", ticketID)
	return NewMessageCardEmbedButton("View", url)
}
