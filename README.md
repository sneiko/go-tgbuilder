# 🤖 go-tgbuilder – Telegram Bot Framework with Menu Builder & FSM

[![Go](https://img.shields.io/badge/Go-1.21+-007d9c?logo=go&logoColor=white)](https://golang.org)
[![License](https://img.shields.io/github/license/sneiko/go-tgbuilder)](https://github.com/sneiko/go-tgbuilder/blob/main/LICENSE)

**go-tgbuilder** is a lightweight, modular Go framework for building Telegram bots with support for dynamic menus, hierarchical navigation, and built-in FSM (Finite State Machine) to manage user states.

Ideal for bots with multi-level menus, FAQs, admin panels, and interactive workflows.

---

## ✨ Features

- 🌐 Tree-based menu system with nested items
- 🔘 Flexible inline and reply keyboard generation
- 🔗 Redirects between menu items
- 🧠 [in_progress] Built-in FSM with pluggable storage (in-memory / Redis) 
- 🔌 Simple `OnClick` event handlers
- 🛠️ Modular, extensible, and easy to test
- 📦 Pure Go, minimal dependencies, no magic

---

## 📦 Installation

```bash
go get github.com/sneiko/go-tgbuilder
```

## 🚀 Quick Start

```go
func main() {
	ctx := context.Background()

	fsm := tgfsm.New(tgfsmdb.NewInMem())

	// Build your menu
	ui := tgbot.NewBuilder(&tgbot.MenuItem{
		ID:      "/start",
		Title:   "Main Menu",
		Message: "Please choose an option:",
		Inline:  true,
		ChildrenRows: []tgbot.MenuItem{
			{
				ID:    "info",
				Row:   0,
				Title: "Information 🚀",
				ChildrenRows: []tgbot.MenuItem{
					{
						ID:         "back",
						Title:      "Back",
						RedirectTo: "/start",
					},
				},
			},
			{
				ID:    "faq",
				Row:   1,
				Title: "Q/A 🚀",
				OnClick: func(ctx context.Context, msg *tgbot.Message) error {
					return msg.SendText("Frequently asked questions...")
				},
			},
		},
	}, nil)

	// Run the bot
	if err := tgbot.NewBot(TgBotToken, false, ui, fsm).Run(ctx); err != nil {
		slog.Error("Failed to run bot", "error", err)
	}
}
```

## 📂 Project Structure 

pkg/
├── tgbot/        - Core bot logic, menu builder, message handling
├── tgfsm/        - Finite State Machine (FSM) manager
└── tgfsmdb/      - Storage implementations: in-memory, Redis


## 🔐 Admin Menu (Optional) 
```go
ui := tgbot.NewBuilder(userMenu, adminMenu)
```

## 💾 State Storage
1. In-Memory (default)
```go
fsm := tgfsm.New(tgfsmdb.NewInMem())
```

2. Redis
```go
redisStorage := tgfsmdb.NewRedis(redis.Options{
	Addr: "localhost:6379",
})
fsm := tgfsm.New(redisStorage)
```   

## 🛡️ Error Handling

All errors are logged via slog. You can customize the logger: 
```go
slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
```

## 🙌 Acknowledgments

Thanks for using this project!
If you like it, please give it a ⭐ and share it with others. 