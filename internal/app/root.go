package app

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spv-dev/chat-client/internal/client"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Чат клиент v0.0.1",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Команда создания",
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Команда отправки",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Команда соединения",
}

var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Создает новый чат",
	Run: func(cmd *cobra.Command, args []string) {
		chatName, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalf("failed to get chat name: %s\n", err.Error())
		}

		ctx := context.Background()
		res, err := client.CreateChat(ctx, chatName)
		if err != nil {
			log.Fatalf("failed to create chat: %s\n", err.Error())
		}

		log.Printf("chat %s created\n", chatName)
		log.Printf("%#v\n", res)

	},
}

var connectChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Соединение с чатом",
	Run: func(cmd *cobra.Command, args []string) {
		chat, err := cmd.Flags().GetString("chat")
		if err != nil {
			log.Fatalf("failed to get chat: %s\n", err.Error())
		}

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatalf("failed to get user: %s\n", err.Error())
		}

		ctx := context.Background()
		res, err := client.ConnectChat(ctx, chat, user)
		if err != nil {
			log.Fatalf("failed to connect chat: %v", err)
		}

		log.Printf("connect to chat %s\n", chat)
		log.Printf("%v\n", res)
	},
}

var sendMessageCmd = &cobra.Command{
	Use:   "message",
	Short: "Отправляет сообщение в чат",
	Run: func(cmd *cobra.Command, args []string) {
		chat, err := cmd.Flags().GetString("chat")
		if err != nil {
			log.Fatalf("failed to get chat: %s\n", err.Error())
		}

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatalf("failed to get user: %s\n", err.Error())
		}

		message, err := cmd.Flags().GetString("body")
		if err != nil {
			log.Fatalf("failed to send message: %s\n", err.Error())
		}

		ctx := context.Background()
		res, err := client.SendMessage(ctx, chat, user, message)
		if err != nil {
			log.Fatalf("failed to connect chat: %v", err)
		}

		log.Printf("send message %s to chat %s\n", message, chat)
		log.Printf("%v\n", res)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(sendCmd)

	createCmd.AddCommand(createChatCmd)
	connectCmd.AddCommand(connectChatCmd)
	sendCmd.AddCommand(sendMessageCmd)

	createChatCmd.Flags().StringP("name", "n", "", "Название чата")
	err := createChatCmd.MarkFlagRequired("name")
	if err != nil {
		log.Fatalf("failed to mark name flag as required: %s\n", err.Error())
	}

	connectChatCmd.Flags().StringP("chat", "c", "", "Идентификатор чата")
	connectChatCmd.Flags().StringP("user", "u", "", "Идентификатор пользователя")
	err = connectChatCmd.MarkFlagRequired("chat")
	if err != nil {
		log.Fatalf("failed to mark chat flag as required: %s\n", err.Error())
	}
	err = connectChatCmd.MarkFlagRequired("user")
	if err != nil {
		log.Fatalf("failed to mark chat flag as required: %s\n", err.Error())
	}

	sendMessageCmd.Flags().StringP("chat", "c", "", "Идентификатор чата")
	sendMessageCmd.Flags().StringP("user", "u", "", "Идентификатор пользователя")
	sendMessageCmd.Flags().StringP("body", "b", "", "Текст сообщения")
	err = sendMessageCmd.MarkFlagRequired("chat")
	if err != nil {
		log.Fatalf("failed to mark chat flag as required: %s\n", err.Error())
	}
	err = sendMessageCmd.MarkFlagRequired("body")
	if err != nil {
		log.Fatalf("failed to mark body flag as required: %s\n", err.Error())
	}
}
