# About GO frameworks 
## ECHO 
The echo framework used in Go is another high-performance, extensible, and minimalist web framework in Golang. It has a highly optimized HTTP router with zero dynamic memory 
allocation that smartly prioritizes routes. It is used to build robust and scalable REST APIs, which can easily be organized into groups. It automatically installs TLS 
certificates from Let’s Encrypt and provides HTTP/2 support which improves the speed and provides a better user experience. It also contains many built-in middlewares to use and 
developers can even define their own which can be set at a root, group, or route level.

It supports data binding for HTTP request payloads, including JSON, XML, or form-data. For data rendering, it contains an API to send a variety of HTTP responses, including JSON, 
XML, HTML, files, and attachments. Templates can be rendered using any template engine and have customized central HTTP error handling. 

The con of using the echo framework is that it is maintained by only a single developer and the code is updated infrequently.
