package loong

import "context"

type userKey struct{}

func (*userKey) String() string {
	return "loong-user-key"
}

var (
    UserKey = &userKey{}

    ContextWithUserHook func(ctx context.Context, u interface{}) context.Context
    UserFromContextHook func(ctx context.Context) interface{}
)
    
func ContextWithUser(ctx context.Context, u interface{}) context.Context {
    if ContextWithUserHook != nil {
        return ContextWithUserHook(ctx, u)
    }
	return context.WithValue(ctx, UserKey, u)
}

func UserFromContext(ctx context.Context) interface{} {
    if UserFromContextHook != nil {
        return UserFromContextHook(ctx)
    }
	return ctx.Value(UserKey)
}

type tokenKey struct{}

func (*tokenKey) String() string {
	return "loong-token-key"
}

var TokenKey = &tokenKey{}

func ContextWithToken(ctx context.Context, u interface{}) context.Context {
	return context.WithValue(ctx, TokenKey, u)
}

func TokenFromContext(ctx context.Context) interface{} {
	return ctx.Value(TokenKey)
}

type sessionKey struct{}

func (*sessionKey) String() string {
	return "loong-session-key"
}

var SessionKey = &sessionKey{}

func ContextWithSession(ctx context.Context, s interface{}) context.Context {
	return context.WithValue(ctx, SessionKey, s)
}

func SessionFromContext(ctx context.Context) interface{} {
	return ctx.Value(SessionKey)
}
