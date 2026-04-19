package schedule

import "time"

func StartHealthCheckTicker(stop <-chan struct{}, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			HealthCheckSchedule()
		case <-stop:
			return
		}
	}
}
