package utils

import (
	"errors"
	"log"
	"strings"
	"sync"
	"time"
)

type CircuitBreakerState int

const (
	CLOSED CircuitBreakerState = iota
	OPEN
	HALF_OPEN
)

type CircuitBreaker struct {
	state        CircuitBreakerState
	failures     int
	lastFailure  time.Time
	mux          sync.Mutex
	failMax      int
	resetTimeout time.Duration
}

func NewCircuitBreaker(failMax int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:        CLOSED,
		failures:     0,
		failMax:      failMax,
		resetTimeout: resetTimeout,
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mux.Lock()
	defer cb.mux.Unlock()

	if cb.state == OPEN {
		if time.Since(cb.lastFailure) > cb.resetTimeout {
			log.Println("Circuit breaker timeout expired. Switching to HALF_OPEN")
			cb.state = HALF_OPEN
		} else {
			return errors.New("circuit is OPEN. Cannot process the request")
		}
	}

	err := fn()
	if err != nil {
		if strings.Contains(err.Error(), "connection error") {
			cb.failures++
			if cb.failures >= cb.failMax {
				cb.state = OPEN
				cb.lastFailure = time.Now()
			}
		}
		return err
	}

	if cb.state == HALF_OPEN {
		log.Println("Circuit breaker is HALF_OPEN. Switching to CLOSED")
		cb.state = CLOSED
		cb.failures = 0
	}

	return nil
}
