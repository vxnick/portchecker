Please note that this is my first Go project, so style and general implementation
is not ideal.

I created this project to check ports are open for gaming purposes.

## Notes

UDP checks are inconsistent: If a port is not open on the server, the client may
show alternating success/failure responses. Consistent success responses will
indicate that the port is actually open.

The server sends a TCP response, but this is not currently checked by the client;
so TCP requests should be considered one-way at present.
