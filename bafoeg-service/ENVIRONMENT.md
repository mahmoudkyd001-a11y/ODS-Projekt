BAFOEGSERVICE_DEBUG   enable debug level for logging
BAFOEGSERVICE_TRACING   enable tracing level for logging
BAFOEGSERVICE_NAME   set the name of the instance of the service
BAFOEGSERVICE_TITLE   set the title in the web page
BAFOEGSERVICE_PORT_NB   the local port of the web service (default=8080)
BAFOEGSERVICE_API_KEYS   space separated list of valid API keys
BAFOEGSERVICE_SESSION_KEY
BAFOEGSERVICE_POLICY   OPA policy for access control
BAFOEGSERVICE_OPASVC   OPA service port to get the OPA policy for access control
BAFOEGSERVICE_REALM   Basic authentication realm
BAFOEGSERVICE_STAFF_USER   username of the administrator
BAFOEGSERVICE_STAFF_PASSWORD   password of the administrator
BAFOEGSERVICE_PARTICIPANT_USER   username of the user
BAFOEGSERVICE_PARTICIPANT_PASSWORD   password of the user
BAFOEGSERVICE_CERT_PEM   certificate for TLS (HTTPS) communication
BAFOEGSERVICE_KEY_PEM   key for TLS (HTTPS) communication
BAFOEGSERVICE_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
BAFOEGSERVICE_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
BAFOEGSERVICE_LOKI_USER   user name as defined for the data source as basic authentication
BAFOEGSERVICE_LOKI_PASSWORD   password as defined for the data source as basic authentication
BAFOEGSERVICE_LOKI_KEY   key/token as defined for the data source
BAFOEGSERVICE_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
BAFOEGSERVICE_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
BAFOEGSERVICE_MAX_DELAY   max time in seconds after which logs are flushed from buffer
BAFOEGSERVICE_LANGUAGE
BAFOEGSERVICE_LANGUAGES
BAFOEGSERVICE_USE_SSE   enable support for _server side event_ communication (default=false)
BAFOEGSERVICE_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
BAFOEGSERVICE_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
BAFOEGSERVICE_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
