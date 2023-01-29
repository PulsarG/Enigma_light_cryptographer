package consts

var NAME_WINDOW_RESULT, // Основа окна
	LABEL_RULES,
	KEY_WORD_TITLE,
	BUTTON_TEXT, // Основная кнопка
	BUTTON_OPEN_FILE,
	TITLE_SUBMENU_ONE, // Меню
	MENU_SUBMENU_ONE_BTN_ONE_TITLE,
	MENU_SUBMENU_ONE_BTN_RUSSIAN,
	MENU_SUBMENU_ONE_BTN_ENGL,
	DIALOG_KEY_WINDOW_TITLE, // Диалог отсутствия ключа-слова
	DIALOG_KEY_ERROR_TEXT,
	DIALOG_KEY_BTN_TITLE,
	BUTTON_RESULTWINDOW_SAVE_FILE, // Окно результата, кнопки
	BUTTON_RESULTWINDOW_COPY,
	DIALOG_ERROR string

const (
	// Версия приложения
	VERSION_APP      = "v0.9.6"
	NAME_WINDOW_MAIN = "Enigma Light " + VERSION_APP

	// Размеры окна
	WINDOW_HEIGHT = 550
	WINDOW_WEIGHT = 550

	// Кнопка сохранения
	BUTTON_SAVE_WEIGHT = 200
	BUTTON_SAVE_HEIGHT = 40
)

func ChangeLang(b bool) {
	if b {
		// Основа окна
		NAME_WINDOW_RESULT = "Результат"
		LABEL_RULES = "Введи текст здесь..."
		KEY_WORD_TITLE = "Ключ-слово"

		// Основная кнопка
		BUTTON_TEXT = "Шифровать / Расшифровать"
		BUTTON_OPEN_FILE = "Открыть файл"

		// Меню
		TITLE_SUBMENU_ONE = "Setting"

		MENU_SUBMENU_ONE_BTN_ONE_TITLE = "Language"

		MENU_SUBMENU_ONE_BTN_RUSSIAN = "Русский"
		MENU_SUBMENU_ONE_BTN_ENGL = "English"

		// Диалог отсутствия ключа-слова
		DIALOG_KEY_WINDOW_TITLE = "Проблема шифрования"
		DIALOG_KEY_ERROR_TEXT = "Не введено ключ-слово"
		DIALOG_KEY_BTN_TITLE = "Ясно"
		/* DIALOG_ERROR = "Ошибка ключа или вводимого текста" */

		// Окно результата, кнопки
		BUTTON_RESULTWINDOW_SAVE_FILE = "Сохранить в файл"
		BUTTON_RESULTWINDOW_COPY = "Скопировать"
	} else {
		// Основа окна
		NAME_WINDOW_RESULT = "Result"
		LABEL_RULES = "Enter text here..."
		KEY_WORD_TITLE = "Key-word"

		// Основная кнопка
		BUTTON_TEXT = "Encrypt / Decrypt"
		BUTTON_OPEN_FILE = "Open file"

		// Меню
		TITLE_SUBMENU_ONE = "Setting"

		MENU_SUBMENU_ONE_BTN_ONE_TITLE = "Language"

		MENU_SUBMENU_ONE_BTN_RUSSIAN = "Русский"
		MENU_SUBMENU_ONE_BTN_ENGL = "English"

		// Диалог отсутствия ключа-слова
		DIALOG_KEY_WINDOW_TITLE = "Encryption problem"
		DIALOG_KEY_ERROR_TEXT = "Keyword not entered"
		DIALOG_KEY_BTN_TITLE = "OK"
	/* 	DIALOG_ERROR = "Key or input text error" */

		// Окно результата, кнопки
		BUTTON_RESULTWINDOW_SAVE_FILE = "Save in file"
		BUTTON_RESULTWINDOW_COPY = "Copy result"
	}
}
