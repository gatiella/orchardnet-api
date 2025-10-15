package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"orchardnet-api/core"
	"os/exec"
	"runtime"
)

const dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>OrchardNet Control Panel</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
        }

        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            max-width: 500px;
            width: 100%;
            padding: 40px;
        }

        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 28px;
        }

        .subtitle {
            color: #666;
            margin-bottom: 30px;
            font-size: 14px;
        }

        .form-group {
            margin-bottom: 20px;
        }

        label {
            display: block;
            color: #555;
            font-weight: 600;
            margin-bottom: 8px;
            font-size: 14px;
        }

        input, select {
            width: 100%;
            padding: 12px 15px;
            border: 2px solid #e0e0e0;
            border-radius: 10px;
            font-size: 14px;
            transition: border-color 0.3s;
        }

        input:focus, select:focus {
            outline: none;
            border-color: #667eea;
        }

        .attack-type {
            display: grid;
            grid-template-columns: repeat(3, 1fr);
            gap: 10px;
            margin-bottom: 20px;
        }

        .attack-btn {
            padding: 15px;
            border: 2px solid #e0e0e0;
            border-radius: 10px;
            background: white;
            cursor: pointer;
            transition: all 0.3s;
            font-weight: 600;
            color: #555;
        }

        .attack-btn:hover {
            border-color: #667eea;
            background: #f8f9ff;
        }

        .attack-btn.active {
            border-color: #667eea;
            background: #667eea;
            color: white;
        }

        .launch-btn {
            width: 100%;
            padding: 15px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 10px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
        }

        .launch-btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 10px 30px rgba(102, 126, 234, 0.4);
        }

        .launch-btn:active {
            transform: translateY(0);
        }

        .response {
            margin-top: 20px;
            padding: 15px;
            border-radius: 10px;
            display: none;
        }

        .response.success {
            background: #d4edda;
            border: 2px solid #c3e6cb;
            color: #155724;
        }

        .response.error {
            background: #f8d7da;
            border: 2px solid #f5c6cb;
            color: #721c24;
        }

        .job-info {
            margin-top: 10px;
            font-family: monospace;
            font-size: 13px;
        }

        .range-value {
            display: inline-block;
            margin-left: 10px;
            color: #667eea;
            font-weight: 600;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🌐 OrchardNet</h1>
        <p class="subtitle">Network Stress Testing Control Panel</p>

        <form id="attackForm">
            <div class="form-group">
                <label>Target Domain/IP</label>
                <input type="text" id="target" placeholder="example.com" required>
            </div>

            <div class="form-group">
                <label>Attack Type</label>
                <div class="attack-type">
                    <button type="button" class="attack-btn active" data-type="http">HTTP</button>
                    <button type="button" class="attack-btn" data-type="syn">SYN</button>
                    <button type="button" class="attack-btn" data-type="udp">UDP</button>
                </div>
            </div>

            <div class="form-group">
                <label>Workers <span class="range-value" id="workersValue">100</span></label>
                <input type="range" id="workers" min="10" max="500" value="100" step="10">
            </div>

            <div class="form-group">
                <label>Duration (seconds) <span class="range-value" id="durationValue">10</span></label>
                <input type="range" id="duration" min="5" max="120" value="10" step="5">
            </div>

            <button type="submit" class="launch-btn">🚀 Launch Attack</button>
        </form>

        <div id="response" class="response"></div>
    </div>

    <script>
        const apiUrl = window.location.origin + '/api/v1/attack';
        const apiKey = 'th0rn3-0rch@rd-k3y-2024';
        let selectedType = 'http';

        document.querySelectorAll('.attack-btn').forEach(btn => {
            btn.addEventListener('click', function() {
                document.querySelectorAll('.attack-btn').forEach(b => b.classList.remove('active'));
                this.classList.add('active');
                selectedType = this.dataset.type;
            });
        });

        document.getElementById('workers').addEventListener('input', function() {
            document.getElementById('workersValue').textContent = this.value;
        });

        document.getElementById('duration').addEventListener('input', function() {
            document.getElementById('durationValue').textContent = this.value;
        });

        document.getElementById('attackForm').addEventListener('submit', async function(e) {
            e.preventDefault();

            const target = document.getElementById('target').value;
            const workers = parseInt(document.getElementById('workers').value);
            const duration = parseInt(document.getElementById('duration').value);

            const payload = {
                target: target,
                type: selectedType,
                workers: workers,
                duration: duration,
                api_key: apiKey
            };

            const responseDiv = document.getElementById('response');
            responseDiv.style.display = 'none';

            try {
                const response = await fetch(apiUrl, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(payload)
                });

                const data = await response.json();

                if (response.ok) {
                    responseDiv.className = 'response success';
                    responseDiv.innerHTML = ` + "`" + `
                        <strong>✓ Attack Launched Successfully!</strong>
                        <div class="job-info">Job ID: ${data.job_id}</div>
                        <div class="job-info">Status: ${data.status}</div>
                        <div class="job-info">Target: ${target}</div>
                        <div class="job-info">Type: ${selectedType.toUpperCase()}</div>
                        <div class="job-info">Duration: ${duration}s</div>
                    ` + "`" + `;
                } else {
                    responseDiv.className = 'response error';
                    responseDiv.innerHTML = ` + "`" + `<strong>✗ Error:</strong> ${data.error || 'Request failed'}` + "`" + `;
                }

                responseDiv.style.display = 'block';

            } catch (error) {
                responseDiv.className = 'response error';
                responseDiv.innerHTML = ` + "`" + `<strong>✗ Connection Error:</strong> ${error.message}` + "`" + `;
                responseDiv.style.display = 'block';
            }
        });
    </script>
</body>
</html>`

func attackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Target   string `json:"target"`
		Type     string `json:"type"`
		Duration int    `json:"duration"`
		Workers  int    `json:"workers"`
		APIKey   string `json:"api_key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.APIKey != "th0rn3-0rch@rd-k3y-2024" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	jobID := core.ScheduleAttack(req.Target, req.Type, req.Workers, req.Duration)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"job_id": jobID,
		"status": "launched",
	})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "completed",
	})
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, dashboardHTML)
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

func main() {
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/api/v1/attack", attackHandler)
	http.HandleFunc("/api/v1/status/", statusHandler)

	port := ":8080"
	url := "http://localhost:8080"

	log.Println("═══════════════════════════════════════")
	log.Println("🌐 OrchardNet API Server Started")
	log.Println("═══════════════════════════════════════")
	log.Printf("📡 Server running on %s\n", port)
	log.Printf("🎯 Dashboard: %s\n", url)
	log.Println("═══════════════════════════════════════")

	// Auto-open browser
	go func() {
		openBrowser(url)
	}()

	log.Fatal(http.ListenAndServe(port, nil))
}
