ANTRAGSERVICE3_DEBUG   enable debug level for logging
ANTRAGSERVICE3_TRACING   enable tracing level for logging
ANTRAGSERVICE3_NAME   set the name of the instance of the service
ANTRAGSERVICE3_TITLE   set the title in the web page
ANTRAGSERVICE3_PORT_NB   the local port of the web service (default=8080)
ANTRAGSERVICE3_API_KEYS   space separated list of valid API keys
ANTRAGSERVICE3_SESSION_KEY
ANTRAGSERVICE3_POLICY   OPA policy for access control
ANTRAGSERVICE3_OPASVC   OPA service port to get the OPA policy for access control
ANTRAGSERVICE3_REALM   Basic authentication realm
ANTRAGSERVICE3_STAFF_USER   username of the administrator
ANTRAGSERVICE3_STAFF_PASSWORD   password of the administrator
ANTRAGSERVICE3_PARTICIPANT_USER   username of the user
ANTRAGSERVICE3_PARTICIPANT_PASSWORD   password of the user
ANTRAGSERVICE3_CERT_PEM   certificate for TLS (HTTPS) communication
ANTRAGSERVICE3_KEY_PEM   key for TLS (HTTPS) communication
ANTRAGSERVICE3_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ANTRAGSERVICE3_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ANTRAGSERVICE3_LOKI_USER   user name as defined for the data source as basic authentication
ANTRAGSERVICE3_LOKI_PASSWORD   password as defined for the data source as basic authentication
ANTRAGSERVICE3_LOKI_KEY   key/token as defined for the data source
ANTRAGSERVICE3_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ANTRAGSERVICE3_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ANTRAGSERVICE3_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ANTRAGSERVICE3_LANGUAGE
ANTRAGSERVICE3_LANGUAGES
ANTRAGSERVICE3_USE_SSE   enable support for _server side event_ communication (default=false)
ANTRAGSERVICE3_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ANTRAGSERVICE3_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ANTRAGSERVICE3_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
