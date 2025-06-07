package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type OTPStore struct {
	client *redis.Client
}

func NewOTPStore(redisURL string) (*OTPStore, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("invalid redis url: %w", err)
	}
	client := redis.NewClient(opts)
	return &OTPStore{client: client}, nil
}

func (o *OTPStore) ValidateSession(ctx context.Context, sessionID string) (bool, error) {
	val, err := o.client.Get(ctx, "otp:session:"+sessionID).Result()
	if err == redis.Nil {
		return false, errors.New("session not found")
	} else if err != nil {
		return false, err
	}

	// Optional: parse TTL or JSON metadata
	if val == "valid" {
		return true, nil
	}
	return false, errors.New("session invalid or expired")
}
