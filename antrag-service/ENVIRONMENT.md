ANTRAGSERVICE_DEBUG   enable debug level for logging
ANTRAGSERVICE_TRACING   enable tracing level for logging
ANTRAGSERVICE_NAME   set the name of the instance of the service
ANTRAGSERVICE_TITLE   set the title in the web page
ANTRAGSERVICE_PORT_NB   the local port of the web service (default=8080)
ANTRAGSERVICE_API_KEYS   space separated list of valid API keys
ANTRAGSERVICE_SESSION_KEY
ANTRAGSERVICE_POLICY   OPA policy for access control
ANTRAGSERVICE_OPASVC   OPA service port to get the OPA policy for access control
ANTRAGSERVICE_REALM   Basic authentication realm
ANTRAGSERVICE_STAFF_USER   username of the administrator
ANTRAGSERVICE_STAFF_PASSWORD   password of the administrator
ANTRAGSERVICE_PARTICIPANT_USER   username of the user
ANTRAGSERVICE_PARTICIPANT_PASSWORD   password of the user
ANTRAGSERVICE_CERT_PEM   certificate for TLS (HTTPS) communication
ANTRAGSERVICE_KEY_PEM   key for TLS (HTTPS) communication
ANTRAGSERVICE_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
ANTRAGSERVICE_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
ANTRAGSERVICE_LOKI_USER   user name as defined for the data source as basic authentication
ANTRAGSERVICE_LOKI_PASSWORD   password as defined for the data source as basic authentication
ANTRAGSERVICE_LOKI_KEY   key/token as defined for the data source
ANTRAGSERVICE_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
ANTRAGSERVICE_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
ANTRAGSERVICE_MAX_DELAY   max time in seconds after which logs are flushed from buffer
ANTRAGSERVICE_LANGUAGE
ANTRAGSERVICE_LANGUAGES
ANTRAGSERVICE_USE_SSE   enable support for _server side event_ communication (default=false)
ANTRAGSERVICE_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
ANTRAGSERVICE_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
ANTRAGSERVICE_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
