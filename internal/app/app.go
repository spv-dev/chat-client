package app

import (
	"context"

	"github.com/spv-dev/chat-client/internal/client"
	"github.com/spv-dev/platform_common/pkg/closer"
)

// App структура приложения
type App struct {
	serviceProvider *serviceProvider
}

// NewApp инициализизует зависимости и создаёт экземпляр структуры приложения
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run запускает приложение
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	client.Init()

	Execute()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
