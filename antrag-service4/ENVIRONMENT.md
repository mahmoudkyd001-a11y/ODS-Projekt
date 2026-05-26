ANTRAGSERVICE4_DEBUG   enable debug level for logging
ANTRAGSERVICE4_TRACING   enable tracing level for logging
ANTRAGSERVICE4_NAME   set the name of the instance of the service
ANTRAGSERVICE4_TITLE   set the title in the web page
ANTRAGSERVICE4_PORT_NB   the local port of the web service (default=8080)
ANTRAGSERVICE4_API_KEYS   space separated list of valid API keys
ANTRAGSERVICE4_SESSION_KEY
ANTRAGSERVICE4_POLICY   OPA policy for access control
ANTRAGSERVICE4_OPASVC   OPA service port to get the OPA policy for access control
ANTRAGSERVICE4_REALM   Basic authentication realm
ANTRAGSERVICE4_STAFF_USER   username of the administrator
ANTRAGSERVICE4_STAFF_PASSWORD   password of the administrator
ANTRAGSERVICE4_PARTICIPANT_USER   username of the user
ANTRAGSERVICE4_PARTICIPANT_PASSWORD   password of the user
ANTRAGSERVICE4_CERT_PEM   certificate for TLS (HTTPS) communication
ANTRAGSERVICE4_KEY_PEM   key for TLS (HTTPS) communication
ANTRAGSERVICE4_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ANTRAGSERVICE4_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ANTRAGSERVICE4_LOKI_USER   user name as defined for the data source as basic authentication
ANTRAGSERVICE4_LOKI_PASSWORD   password as defined for the data source as basic authentication
ANTRAGSERVICE4_LOKI_KEY   key/token as defined for the data source
ANTRAGSERVICE4_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ANTRAGSERVICE4_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ANTRAGSERVICE4_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ANTRAGSERVICE4_LANGUAGE
ANTRAGSERVICE4_LANGUAGES
ANTRAGSERVICE4_USE_SSE   enable support for _server side event_ communication (default=false)
ANTRAGSERVICE4_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ANTRAGSERVICE4_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ANTRAGSERVICE4_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
