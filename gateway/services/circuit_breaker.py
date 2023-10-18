import logging
import time
from functools import wraps
from enum import Enum

import grpc
from fastapi import HTTPException


class CircuitBreakerState(Enum):
    CLOSED = "CLOSED"
    OPEN = "OPEN"
    HALF_OPEN = "HALF_OPEN"


class CircuitBreaker:
    def __init__(self, fail_max, reset_timeout):
        self.fail_max = fail_max
        self.reset_timeout = reset_timeout
        self.state = CircuitBreakerState.CLOSED
        self.failures = 0
        self.last_failure_time = None
        self.successful_calls = 0
        self.success_threshold = fail_max
        self.logger = logging.getLogger("Started cb: " + __name__)

    def transition(self, new_state):
        self.state = new_state
        if new_state == CircuitBreakerState.OPEN:
            self.last_failure_time = time.time()
        self.logger.info(f"Circuit transitioned to {new_state}. Failures: {self.failures}")

    def __call__(self, func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            if self.state == CircuitBreakerState.OPEN:
                if time.time() - self.last_failure_time > self.reset_timeout:
                    self.transition(CircuitBreakerState.HALF_OPEN)
                else:
                    raise Exception("Circuit is OPEN. Cannot process the request.")

            try:
                result = func(*args, **kwargs)
                if self.state == CircuitBreakerState.HALF_OPEN:
                    self.successful_calls += 1
                    if self.successful_calls >= self.success_threshold:
                        self.successful_calls = 0
                        self.transition(CircuitBreakerState.CLOSED)
                return result
            except HTTPException as e:
                if "DNS resolution failed" in e.detail or "failed to connect to all addresses" in e.detail:
                    self.failures += 1
                    if self.failures >= self.fail_max:
                        self.transition(CircuitBreakerState.OPEN)
                raise e
            except grpc.RpcError as e:
                if e.code() in [grpc.StatusCode.UNAVAILABLE, grpc.StatusCode.DEADLINE_EXCEEDED]:
                    self.failures += 1
                    if self.failures >= self.fail_max:
                        self.transition(CircuitBreakerState.OPEN)
                raise e

        return wrapper
