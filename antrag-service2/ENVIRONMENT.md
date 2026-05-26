ANTRAGSERVICE2_DEBUG   enable debug level for logging
ANTRAGSERVICE2_TRACING   enable tracing level for logging
ANTRAGSERVICE2_NAME   set the name of the instance of the service
ANTRAGSERVICE2_TITLE   set the title in the web page
ANTRAGSERVICE2_PORT_NB   the local port of the web service (default=8080)
ANTRAGSERVICE2_API_KEYS   space separated list of valid API keys
ANTRAGSERVICE2_SESSION_KEY
ANTRAGSERVICE2_POLICY   OPA policy for access control
ANTRAGSERVICE2_OPASVC   OPA service port to get the OPA policy for access control
ANTRAGSERVICE2_REALM   Basic authentication realm
ANTRAGSERVICE2_STAFF_USER   username of the administrator
ANTRAGSERVICE2_STAFF_PASSWORD   password of the administrator
ANTRAGSERVICE2_PARTICIPANT_USER   username of the user
ANTRAGSERVICE2_PARTICIPANT_PASSWORD   password of the user
ANTRAGSERVICE2_CERT_PEM   certificate for TLS (HTTPS) communication
ANTRAGSERVICE2_KEY_PEM   key for TLS (HTTPS) communication
ANTRAGSERVICE2_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ANTRAGSERVICE2_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ANTRAGSERVICE2_LOKI_USER   user name as defined for the data source as basic authentication
ANTRAGSERVICE2_LOKI_PASSWORD   password as defined for the data source as basic authentication
ANTRAGSERVICE2_LOKI_KEY   key/token as defined for the data source
ANTRAGSERVICE2_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ANTRAGSERVICE2_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ANTRAGSERVICE2_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ANTRAGSERVICE2_LANGUAGE
ANTRAGSERVICE2_LANGUAGES
ANTRAGSERVICE2_USE_SSE   enable support for _server side event_ communication (default=false)
ANTRAGSERVICE2_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ANTRAGSERVICE2_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ANTRAGSERVICE2_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
