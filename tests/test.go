package tests

import "context"

type WeChatService struct {
}

// Deprecated: V1.4.0 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
func (svc WeChatService) GetWxaCodeUnlimited(ctx context.Context) ([]byte, error) {
	panic("implement me")
}

func a() {
	WeChatService{}.GetWxaCodeUnlimited(context.Background())
}

// Deprecated: V1.4.0 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
type X struct {
	// Deprecated: 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
	ABD string
}

// Deprecated: V1.4.0 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
type XXX interface {
	// Deprecated: 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
	ABD()
}

var (
	// Deprecated: 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
	ax int = 0
)

// Deprecated: V1.4.0 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
var abc = 0

// Deprecated: V1.4.0 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
var abcd = 0

const (
	// Deprecated: 请使用 GetWxaCodeUnlimitedV1, 该方法将会在 v1.4.0 版本中移除
	AX = 0
)
