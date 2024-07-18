wrk.method = "POST"
wrk.body   = '{"talk_name": "Go e Microserviços", "score": "10"}'
wrk.headers["Authorization"] = os.getenv("TOKEN")
