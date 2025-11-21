package domain

import "errors"

var (
	ErrNotFound          = errors.New("объект не найден")
	ErrUserNotFound      = errors.New("пользователь не найден")
	ErrTeamNotFound      = errors.New("команда не найдена")
	ErrPRNotFound        = errors.New("пул реквест не найден")
	ErrUserAlreadyExists = errors.New("пользователь уже существует")
	ErrTeamAlreadyExists = errors.New("команда уже существует")
	ErrPRAlreadyExists   = errors.New("пул реквест уже существует")
	ErrInvalidInput      = errors.New("не валидный ввод данных")
	ErrAccessDenied      = errors.New("доступ запрещен")
)
