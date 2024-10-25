package domain

import (
	"my_app/src/infrastructure"
)

type UserPorter interface {
	infrastructure.Porter[any, User]
	// tu jakieś nowe w razie czego
}
