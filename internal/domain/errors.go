package domain

import "errors"

var (
	ErrNotFound     = errors.New("объект не найден")
	ErrInvalidInput = errors.New("не валидный ввод данных")
	ErrAccessDenied = errors.New("доступ запрещен")

	ErrUserNotFound      = errors.New("пользователь не найден")
	ErrUserAlreadyExists = errors.New("пользователь уже существует")
	ErrUserNotActive     = errors.New("пользователь не активен")

	ErrTeamNotFound      = errors.New("команда не найдена")
	ErrTeamAlreadyExists = errors.New("команда уже существует")

	ErrPRNotFound           = errors.New("пул реквест не найден")
	ErrPRAlreadyExists      = errors.New("пул реквест уже существует")
	ErrPRAlreadyMerged      = errors.New("пул реквест не уже замерджен")
	ErrNoReviewersAvailable = errors.New("нет активных ревьюверов")
)
