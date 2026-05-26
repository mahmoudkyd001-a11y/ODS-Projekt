ANTRAGSERVICE5_DEBUG   enable debug level for logging
ANTRAGSERVICE5_TRACING   enable tracing level for logging
ANTRAGSERVICE5_NAME   set the name of the instance of the service
ANTRAGSERVICE5_TITLE   set the title in the web page
ANTRAGSERVICE5_PORT_NB   the local port of the web service (default=8080)
ANTRAGSERVICE5_API_KEYS   space separated list of valid API keys
ANTRAGSERVICE5_SESSION_KEY
ANTRAGSERVICE5_POLICY   OPA policy for access control
ANTRAGSERVICE5_OPASVC   OPA service port to get the OPA policy for access control
ANTRAGSERVICE5_REALM   Basic authentication realm
ANTRAGSERVICE5_STAFF_USER   username of the administrator
ANTRAGSERVICE5_STAFF_PASSWORD   password of the administrator
ANTRAGSERVICE5_PARTICIPANT_USER   username of the user
ANTRAGSERVICE5_PARTICIPANT_PASSWORD   password of the user
ANTRAGSERVICE5_CERT_PEM   certificate for TLS (HTTPS) communication
ANTRAGSERVICE5_KEY_PEM   key for TLS (HTTPS) communication
ANTRAGSERVICE5_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ANTRAGSERVICE5_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ANTRAGSERVICE5_LOKI_USER   user name as defined for the data source as basic authentication
ANTRAGSERVICE5_LOKI_PASSWORD   password as defined for the data source as basic authentication
ANTRAGSERVICE5_LOKI_KEY   key/token as defined for the data source
ANTRAGSERVICE5_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ANTRAGSERVICE5_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ANTRAGSERVICE5_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ANTRAGSERVICE5_LANGUAGE
ANTRAGSERVICE5_LANGUAGES
ANTRAGSERVICE5_USE_SSE   enable support for _server side event_ communication (default=false)
ANTRAGSERVICE5_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ANTRAGSERVICE5_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ANTRAGSERVICE5_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
