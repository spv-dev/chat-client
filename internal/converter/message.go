package converter

import (
	"time"

	desc "github.com/spv-dev/chat-server/pkg/chat_v1"

	"github.com/spv-dev/chat-client/internal/model"
)

// ToMessageInfoFromDesc конвертер MessageInfo из API слоя в сервисный слой
func ToMessageFromAPI(mess *desc.Message) *model.Message {
	if mess == nil {
		return &model.Message{}
	}
	var updatedAt time.Time
	if mess.UpdatedAt != nil {
		updatedAt = mess.UpdatedAt.AsTime()
	}
	var deletedAt time.Time
	if mess.DeletedAt != nil {
		deletedAt = mess.DeletedAt.AsTime()
	}

	return &model.Message{
		ID: mess.Id,
		Info: model.MessageInfo{
			ChatID: mess.Info.ChatId,
			UserID: mess.Info.UserId,
			Body:   mess.Info.Body,
		},
		State:     mess.State,
		Type:      mess.Type,
		CreatedAt: mess.CreatedAt.AsTime(),
		UpdatedAt: &updatedAt,
		DeletedAt: &deletedAt,
	}
}
