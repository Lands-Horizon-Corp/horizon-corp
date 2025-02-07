package middleware

import (
	"time"

	"go.uber.org/zap"
)

func (m *Middleware) alertAdmin(ip string, path string) {
	m.logger.Warn("Admin Alert: Suspicious activity detected", zap.String("ip", ip), zap.String("path", path))
}

func (m *Middleware) setWithRetry(key string, value string, duration time.Duration, retries int) error {
	var err error
	for i := 0; i < retries; i++ {
		err = m.cache.Set(key, value, duration)
		if err == nil {
			return nil
		}
		m.logger.Warn("Cache set failed, retrying...", zap.String("key", key), zap.Int("attempt", i+1), zap.Error(err))
		time.Sleep(time.Second * 2)
	}
	return err
}
