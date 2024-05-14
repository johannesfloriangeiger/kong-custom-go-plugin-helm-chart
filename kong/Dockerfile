FROM kong

COPY add-header/add-header /usr/local/kong/add-header
ENV KONG_PLUGINS="bundled,add-header"
ENV KONG_PLUGINSERVER_NAMES="add-header"
ENV KONG_PLUGINSERVER_ADD_HEADER_START_CMD="/usr/local/kong/add-header"
ENV KONG_PLUGINSERVER_ADD_HEADER_QUERY_CMD="/usr/local/kong/add-header -dump"

COPY kong.yaml /

RUN cp /etc/kong/kong.conf.default /etc/kong/kong.conf