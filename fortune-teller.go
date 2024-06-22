package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
// telegram token
const TOKEN = "My token"

var bot *tgbotapi.BotAPI
var chatID int64

func connectWithTelegram() {
	// будем получать новый апи
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot conect to Telegram")
	}
}

var fortuneTellerNames = [3]string{"стивен", "стив", "кинг"}

var answers = []string{
	"Да",
	"Нет",
	"Большинство страхов ничем не обоснованы и никогда не сбываются. (Сияние)",
	"Никогда не останавливайся на дороге, где погибли более трех людей. (Оно)",
	"Смерть - это ребенок, который спит. (Бессонница)",
	"Никогда не стоит узнавать, что находится в темноте. (Туман)",
	"Не верьте глазам - они обманчивы. (Заключенная)",
	"Иногда зло начинается с того, что кто-то хочет быть добрым. (Сияние)",
	"Никогда не поднимайте палку на собаку, которая может съесть вас по кусочкам. (Куджо)",
	"Чтобы узнать, кто вы есть на самом деле, найдите способ выжить в самом темном месте. (Темная башня)",
	"Последние слова не означают ничего. Это то, что происходит после, важно. (Зона смерти)",
	"Если вы намерены что-то убить, убедитесь, что это действительно мертво. (Кладбище домашних животных)",
}

func sandMassage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatID, msg)
	bot.Send(msgConfig)
}

func isMessegeForfortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}
	msgInLoverCase := strings.ToLower(update.Message.Text)
	for _, name := range fortuneTellerNames {
		if strings.Contains(msgInLoverCase, name) {
			return true
		}
	}
	return false
}

func getFortuneTellerAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatID, getFortuneTellerAnswer())
	msg.ReplyToMessageID = int(update.Message.Chat.ID)
	sandMassage(msg.Text)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatID = update.Message.Chat.ID
			sandMassage("Задай свой вопрос, назвав меня по имени.Ответом на вопрос должно быть \"Да\" либо \"Нет\", Например, \"Кинг," +
				" Я готов сменить работу\" либо\"Стивен, я действительно хочу отправиться на эту вечеринку?\"")
		}
		if isMessegeForfortuneTeller(&update) {
			sendAnswer(&update)
		}
	}
}
