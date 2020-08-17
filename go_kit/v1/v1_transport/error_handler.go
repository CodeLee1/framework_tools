package v1_transport

import (
	"context"
	"fmt"
	"framework_tools/go_kit/v1/v1_service"
	"go.uber.org/zap"
)

// transport传输错误处理，当前只是日志记录

type LogErrorHandler struct {
	logger *zap.Logger
}

func NewZapLogErrorHandler(logger *zap.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(v1_service.ContextReqUUid)), zap.Error(err))
}
