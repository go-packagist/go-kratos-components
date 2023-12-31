package env

import "context"

func Provider(env Env) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		SetEnv(env)

		return nil
	}
}
