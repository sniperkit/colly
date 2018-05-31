# LogCounter
A simple and lightweight set of utilities to give overviews of logs following logfmt records. 

Keeps a configurable amount of Top N logs per log level, using a generally constant amount of memory.

## File Structure
`display`: A CLI application giving a view of the top logs. Connects to the reader server to get up to date information.

`reader`: Your server's logs are piped through the `reader`, which spins up a websocket server to push updates out to any clients. 

`random`: Helper package to generate random logs. 

`printer`: Simple program to spit out tons of random logs.
