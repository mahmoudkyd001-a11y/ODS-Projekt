ANTRAGSERVICE6_DEBUG   enable debug level for logging
ANTRAGSERVICE6_TRACING   enable tracing level for logging
ANTRAGSERVICE6_NAME   set the name of the instance of the service
ANTRAGSERVICE6_TITLE   set the title in the web page
ANTRAGSERVICE6_PORT_NB   the local port of the web service (default=8080)
ANTRAGSERVICE6_API_KEYS   space separated list of valid API keys
ANTRAGSERVICE6_SESSION_KEY
ANTRAGSERVICE6_POLICY   OPA policy for access control
ANTRAGSERVICE6_OPASVC   OPA service port to get the OPA policy for access control
ANTRAGSERVICE6_REALM   Basic authentication realm
ANTRAGSERVICE6_STAFF_USER   username of the administrator
ANTRAGSERVICE6_STAFF_PASSWORD   password of the administrator
ANTRAGSERVICE6_PARTICIPANT_USER   username of the user
ANTRAGSERVICE6_PARTICIPANT_PASSWORD   password of the user
ANTRAGSERVICE6_CERT_PEM   certificate for TLS (HTTPS) communication
ANTRAGSERVICE6_KEY_PEM   key for TLS (HTTPS) communication
ANTRAGSERVICE6_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ANTRAGSERVICE6_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ANTRAGSERVICE6_LOKI_USER   user name as defined for the data source as basic authentication
ANTRAGSERVICE6_LOKI_PASSWORD   password as defined for the data source as basic authentication
ANTRAGSERVICE6_LOKI_KEY   key/token as defined for the data source
ANTRAGSERVICE6_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ANTRAGSERVICE6_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ANTRAGSERVICE6_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ANTRAGSERVICE6_LANGUAGE
ANTRAGSERVICE6_LANGUAGES
ANTRAGSERVICE6_USE_SSE   enable support for _server side event_ communication (default=false)
ANTRAGSERVICE6_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ANTRAGSERVICE6_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ANTRAGSERVICE6_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
