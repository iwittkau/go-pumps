Go-Pumps
===
> simple channel management

I was working on a lot of channels recently. I saw that I was always using the same techniques to write and read from go channels. So I made a small library which will force me to be more consistence.

This library hides all channel related handling. So you could use channels without actually having to care about [The Channel Closing Principle](http://www.tapirgames.com/blog/golang-channel-closing).

## Getting started

```
go get -u github.com/iwittkau/go-pumps
```

See `cmd/example` for an example

## Features

TBA

## Contributing

As I use this for my own projects, I know this might not be the perfect approach
for all the projects out there. If you have any ideas, just
[open an issue][issues] and tell me what you think.

## Licensing

This project is licensed under the MIT License.

# Shout outs

- Thanks to [jehna](https://github.com/jehna) for the [README boilerplate](https://github.com/jehna/readme-best-practices).

[issues]:https://github.com/iwittkau/go-pumps/issues/new
