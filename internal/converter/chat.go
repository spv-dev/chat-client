package converter

import (
	"time"

	desc "github.com/spv-dev/chat-server/pkg/chat_v1"

	"github.com/spv-dev/chat-client/internal/model"
)

// ToChatFromModel конвертер Chat в API слой из сервисного слоя
func ToChatFromAPI(chat *desc.Chat) *model.Chat {
	if chat == nil {
		return &model.Chat{}
	}
	var updatedAt time.Time
	if chat.UpdatedAt != nil {
		updatedAt = chat.UpdatedAt.AsTime()
	}
	var deletedAt time.Time
	if chat.DeletedAt != nil {
		deletedAt = chat.DeletedAt.AsTime()
	}

	return &model.Chat{
		ID:    chat.Id,
		State: chat.State,
		Info: model.ChatInfo{
			Title: chat.Info.Title,
		},
		CreatedAt: chat.CreatedAt.AsTime(),
		UpdatedAt: &updatedAt,
		DeletedAt: &deletedAt,
	}
}
