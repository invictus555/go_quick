package structs

import (
	"context"
	"fmt"

	"code.byted.org/bytedance/app_determination"
)

func TestifyTikTokAppDetermmination() {
	ctx := context.TODO()
	appDeterminationClient := app_determination.MustNewAppDeterminationClient(ctx)

	for i := 0; i < 10000; i++ {
		if appDeterminationClient.IsTiktokAppId(ctx, int64(i)) {
			fmt.Printf("Is TikTok App:", i)
		} else {
			fmt.Printf("Is Not TikTok App:", i)
		}
	}
}
