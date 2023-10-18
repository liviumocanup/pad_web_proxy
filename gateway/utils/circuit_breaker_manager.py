from services.circuit_breaker import CircuitBreaker


class CircuitBreakerManager:
    _circuit_breakers = {}

    @classmethod
    def get_breaker(cls, service_name, fail_max=3, reset_timeout=10):
        """Get (or create) a circuit breaker for a specific service."""
        if service_name not in cls._circuit_breakers:
            cls._circuit_breakers[service_name] = CircuitBreaker(fail_max, reset_timeout)
        return cls._circuit_breakers[service_name]
