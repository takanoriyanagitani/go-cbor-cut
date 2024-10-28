package cborcut

import (
	"context"
)

type Cut func(context.Context) error
