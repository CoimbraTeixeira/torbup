#
# Used to download file from EDGE and ORIGIN servers to compare speed, size, error, etc... 
# Create to check if the CDN delivery faster then the origin server.
#
/usr/bin/curl -so /dev/null -H "Pragma: no-cache" -H "Cache-Control:no-cache" -w \
"\tHTTP Code: %{http_code} \
Content Type: %{content_type} \
\tHTTP Connect: %{http_connect} \
\tRedirects: %{num_redirects} \
\tTotal Time: %{time_total} \
\tSize:%{size_download} \
\tSpeed  %{speed_download} \
\n" "$URL?`date +%s`"