BAUM_SERVICE_DEBUG   enable debug level for logging
BAUM_SERVICE_TRACING   enable tracing level for logging
BAUM_SERVICE_NAME   set the name of the instance of the service
BAUM_SERVICE_TITLE   set the title in the web page
BAUM_SERVICE_PORT_NB   the local port of the web service (default=8080)
BAUM_SERVICE_API_KEYS   space separated list of valid API keys
BAUM_SERVICE_SESSION_KEY
BAUM_SERVICE_POLICY   OPA policy for access control
BAUM_SERVICE_OPASVC   OPA service port to get the OPA policy for access control
BAUM_SERVICE_REALM   Basic authentication realm
BAUM_SERVICE_STAFF_USER   username of the administrator
BAUM_SERVICE_STAFF_PASSWORD   password of the administrator
BAUM_SERVICE_PARTICIPANT_USER   username of the user
BAUM_SERVICE_PARTICIPANT_PASSWORD   password of the user
BAUM_SERVICE_CERT_PEM   certificate for TLS (HTTPS) communication
BAUM_SERVICE_KEY_PEM   key for TLS (HTTPS) communication
BAUM_SERVICE_LOG_FILE   filename of the logging file or if it is "-" log all messages to the console
BAUM_SERVICE_LOKI_SERVER   URL of the Loki Server, e.g. the Grafana Cloud
BAUM_SERVICE_LOKI_USER   user name as defined for the data source as basic authentication
BAUM_SERVICE_LOKI_PASSWORD   password as defined for the data source as basic authentication
BAUM_SERVICE_LOKI_KEY   key/token as defined for the data source
BAUM_SERVICE_LABELS   set of labels for Grafana Loki, like "app:myapp,tenant:mycustomer"
BAUM_SERVICE_BUFFERSIZE   number of log lines cached before sending to Grafana Loki
BAUM_SERVICE_MAX_DELAY   max time in seconds after which logs are flushed from buffer
BAUM_SERVICE_LANGUAGE
BAUM_SERVICE_LANGUAGES
BAUM_SERVICE_USE_SSE   enable support for _server side event_ communication (default=false)
BAUM_SERVICE_PROGRESS_DURATION   default duration of the progress bar (default=100ms)
BAUM_SERVICE_RAPIDOC_DOC   enable Rapidoc for the OpenAPI viewer (default=false)
BAUM_SERVICE_ELEMENTS_DOC   enable Elements for the OpenAPI viewer (default=false)
