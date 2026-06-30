VACATIONSERVICE_DEBUG   enable debug level for logging
VACATIONSERVICE_TRACING   enable tracing level for logging
VACATIONSERVICE_NAME   set the name of the instance of the service
VACATIONSERVICE_TITLE   set the title in the web page
VACATIONSERVICE_PORT_NB   the local port of the web service (default=8080)
VACATIONSERVICE_API_KEYS   space separated list of valid API keys
VACATIONSERVICE_SESSION_KEY
VACATIONSERVICE_POLICY   OPA policy for access control
VACATIONSERVICE_OPASVC   OPA service port to get the OPA policy for access control
VACATIONSERVICE_REALM   Basic authentication realm
VACATIONSERVICE_STAFF_USER   username of the administrator
VACATIONSERVICE_STAFF_PASSWORD   password of the administrator
VACATIONSERVICE_PARTICIPANT_USER   username of the user
VACATIONSERVICE_PARTICIPANT_PASSWORD   password of the user
VACATIONSERVICE_CERT_PEM   certificate for TLS (HTTPS) communication
VACATIONSERVICE_KEY_PEM   key for TLS (HTTPS) communication
VACATIONSERVICE_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
VACATIONSERVICE_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
VACATIONSERVICE_LOKI_USER   user name as defined for the data source as basic authentication
VACATIONSERVICE_LOKI_PASSWORD   password as defined for the data source as basic authentication
VACATIONSERVICE_LOKI_KEY   key/token as defined for the data source
VACATIONSERVICE_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
VACATIONSERVICE_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
VACATIONSERVICE_MAX_DELAY   max time in seconds after which logs are flushed from buffer
VACATIONSERVICE_LANGUAGE
VACATIONSERVICE_LANGUAGES
VACATIONSERVICE_USE_SSE   enable support for _server side event_ communication (default=false)
VACATIONSERVICE_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
VACATIONSERVICE_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
VACATIONSERVICE_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
