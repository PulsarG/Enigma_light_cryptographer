package consts

var NAME_WINDOW_MAIN, // Основа окна
	NAME_WINDOW_RESULT,
	LABEL_RULES,
	KEY_WORD_TITLE,
	BUTTON_TEXT, // Основная кнопка
	BUTTON_OPEN_FILE,
	TITLE_SUBMENU_ONE, // Меню
	MENU_SUBMENU_ONE_BTN_ONE_TITLE,
	MENU_SUBMENU_ONE_BTN_RUSSIAN,
	MENU_SUBMENU_ONE_BTN_ENGL,
	DIALOG_KEY_WINDOW_TITILE, // Диалог отсутствия ключа-слова
	DIALOG_KEY_ERROR_TEXT,
	DIALOG_KEY_BTN_TITLE,
	BUTTON_RESULTWINDOW_SAVE_FILE, // Окно результата, кнопки
	BUTTON_RESULTWINDOW_COPY string

var WINDOW_HEIGHT, // Размеры окна
	WINDOW_WEIGHT,
	BUTTON_SAVE_WEIGHT, // Кнопка сохранения
	BUTTON_SAVE_HEIGHT float32

func ChangeLang(b bool) {
	if b {
		// Основа окна
		NAME_WINDOW_MAIN = "Enigma Light v0.9.4"
		NAME_WINDOW_RESULT = "Результат"
		LABEL_RULES = "Только буквы русского алфавита нижнего регистра," + "\n" + "пробел и символы , . ! ? - ( )"
		KEY_WORD_TITLE = "Ключ-слово"

		// Основная кнопка
		BUTTON_TEXT = "Шифровать / Расшифровать"
		BUTTON_OPEN_FILE = "Открыть файл"

		// Размеры окна
		WINDOW_HEIGHT = 550
		WINDOW_WEIGHT = 550

		// Кнопка сохранения
		BUTTON_SAVE_WEIGHT = 200
		BUTTON_SAVE_HEIGHT = 40

		// Меню
		TITLE_SUBMENU_ONE = "Setting"

		MENU_SUBMENU_ONE_BTN_ONE_TITLE = "Language"

		MENU_SUBMENU_ONE_BTN_RUSSIAN = "Русский"
		MENU_SUBMENU_ONE_BTN_ENGL = "English"

		// Диалог отсутствия ключа-слова
		DIALOG_KEY_WINDOW_TITILE = "Проблема шифрования"
		DIALOG_KEY_ERROR_TEXT = "Не введено ключ-слово"
		DIALOG_KEY_BTN_TITLE = "Ясно"

		// Окно результата, кнопки
		BUTTON_RESULTWINDOW_SAVE_FILE = "Сохранить в файл"
		BUTTON_RESULTWINDOW_COPY = "Скопировать"
	} else {
		// Основа окна
		NAME_WINDOW_MAIN = "Enigma Light v0.9.4"
		NAME_WINDOW_RESULT = "Result"
		LABEL_RULES = "Only lowercase letters of the Russian alphabet, space and symbols" + "\n" + " , . ! ? - ( )"
		KEY_WORD_TITLE = "Key-word"

		// Основная кнопка
		BUTTON_TEXT = "Crypt / Uncrypt"
		BUTTON_OPEN_FILE = "Open file"

		// Размеры окна
		WINDOW_HEIGHT = 550
		WINDOW_WEIGHT = 550

		// Кнопка сохранения
		BUTTON_SAVE_WEIGHT = 200
		BUTTON_SAVE_HEIGHT = 40

		// Меню
		TITLE_SUBMENU_ONE = "Setting"

		MENU_SUBMENU_ONE_BTN_ONE_TITLE = "Language"

		MENU_SUBMENU_ONE_BTN_RUSSIAN = "Русский"
		MENU_SUBMENU_ONE_BTN_ENGL = "English"

		// Диалог отсутствия ключа-слова
		DIALOG_KEY_WINDOW_TITILE = "Проблема шифрования"
		DIALOG_KEY_ERROR_TEXT = "Не введено ключ-слово"
		DIALOG_KEY_BTN_TITLE = "Ясно"

		// Окно результата, кнопки
		BUTTON_RESULTWINDOW_SAVE_FILE = "Save in file"
		BUTTON_RESULTWINDOW_COPY = "Copy result"
	}
}
