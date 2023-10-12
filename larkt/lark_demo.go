package larkt

import (
	"fmt"
	"time"
)

func TestifyLarkMessage() {
	ticket := &DBTicket{
		ID:           1,
		Creator:      "shengchao",
		Reviewer:     "shengchao",
		Description:  "this is first test",
		CreatedAt:    time.Now(),
		ReviewStatus: "pending",
	}

	InitLarkRobot()

	msg, _ := LarkNotificationMessageCard4RequestReview(ticket)
	fmt.Println(msg)
	LarkRobot.SendLarkCardMessageToReceiver(ticket.Reviewer, msg)
}
