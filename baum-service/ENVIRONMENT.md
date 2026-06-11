BAUMSERVICE_DEBUG   enable debug level for logging
BAUMSERVICE_TRACING   enable tracing level for logging
BAUMSERVICE_NAME   set the name of the instance of the service
BAUMSERVICE_TITLE   set the title in the web page
BAUMSERVICE_PORT_NB   the local port of the web service (default=8080)
BAUMSERVICE_API_KEYS   space separated list of valid API keys
BAUMSERVICE_SESSION_KEY
BAUMSERVICE_POLICY   OPA policy for access control
BAUMSERVICE_OPASVC   OPA service port to get the OPA policy for access control
BAUMSERVICE_REALM   Basic authentication realm
BAUMSERVICE_STAFF_USER   username of the administrator
BAUMSERVICE_STAFF_PASSWORD   password of the administrator
BAUMSERVICE_PARTICIPANT_USER   username of the user
BAUMSERVICE_PARTICIPANT_PASSWORD   password of the user
BAUMSERVICE_CERT_PEM   certificate for TLS (HTTPS) communication
BAUMSERVICE_KEY_PEM   key for TLS (HTTPS) communication
BAUMSERVICE_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
BAUMSERVICE_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
BAUMSERVICE_LOKI_USER   user name as defined for the data source as basic authentication
BAUMSERVICE_LOKI_PASSWORD   password as defined for the data source as basic authentication
BAUMSERVICE_LOKI_KEY   key/token as defined for the data source
BAUMSERVICE_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
BAUMSERVICE_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
BAUMSERVICE_MAX_DELAY   max time in seconds after which logs are flushed from buffer
BAUMSERVICE_LANGUAGE
BAUMSERVICE_LANGUAGES
BAUMSERVICE_USE_SSE   enable support for _server side event_ communication (default=false)
BAUMSERVICE_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
BAUMSERVICE_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
BAUMSERVICE_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
