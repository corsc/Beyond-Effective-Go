package loader

import (
	"context"

	"github.com/corsc/Beyond-Effective-Go/Chapter09/03_code_generation/02_stub_generator/01_go_ast/user"
)

type UserLoader interface {
	LoadByID(ctx context.Context, userID int64) (*user.User, error)
}
