ANTRAGSERVICE7_DEBUG   enable debug level for logging
ANTRAGSERVICE7_TRACING   enable tracing level for logging
ANTRAGSERVICE7_NAME   set the name of the instance of the service
ANTRAGSERVICE7_TITLE   set the title in the web page
ANTRAGSERVICE7_PORT_NB   the local port of the web service (default=8080)
ANTRAGSERVICE7_API_KEYS   space separated list of valid API keys
ANTRAGSERVICE7_SESSION_KEY
ANTRAGSERVICE7_POLICY   OPA policy for access control
ANTRAGSERVICE7_OPASVC   OPA service port to get the OPA policy for access control
ANTRAGSERVICE7_REALM   Basic authentication realm
ANTRAGSERVICE7_STAFF_USER   username of the administrator
ANTRAGSERVICE7_STAFF_PASSWORD   password of the administrator
ANTRAGSERVICE7_PARTICIPANT_USER   username of the user
ANTRAGSERVICE7_PARTICIPANT_PASSWORD   password of the user
ANTRAGSERVICE7_CERT_PEM   certificate for TLS (HTTPS) communication
ANTRAGSERVICE7_KEY_PEM   key for TLS (HTTPS) communication
ANTRAGSERVICE7_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ANTRAGSERVICE7_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ANTRAGSERVICE7_LOKI_USER   user name as defined for the data source as basic authentication
ANTRAGSERVICE7_LOKI_PASSWORD   password as defined for the data source as basic authentication
ANTRAGSERVICE7_LOKI_KEY   key/token as defined for the data source
ANTRAGSERVICE7_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ANTRAGSERVICE7_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ANTRAGSERVICE7_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ANTRAGSERVICE7_LANGUAGE
ANTRAGSERVICE7_LANGUAGES
ANTRAGSERVICE7_USE_SSE   enable support for _server side event_ communication (default=false)
ANTRAGSERVICE7_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ANTRAGSERVICE7_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ANTRAGSERVICE7_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
