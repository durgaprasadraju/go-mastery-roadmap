package solutions

// HealthStatus represents probe result.
type HealthStatus struct {
	OK     bool
	Reason string
}

// Liveness returns alive status.
func Liveness() HealthStatus { return HealthStatus{OK: true} }

// Readiness checks dependencies (stub).
func Readiness(dbOK, cacheOK bool) HealthStatus {
	if dbOK && cacheOK {
		return HealthStatus{OK: true}
	}
	return HealthStatus{OK: false, Reason: "dependency unavailable"}
}