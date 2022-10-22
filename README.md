# rpiSystemMonitor

<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Stars][stars-shield]][stars-url]
[![Forks][forks-shield]][forks-url]
[![Sourcegraph][sourcegraph-shield]][sourcegraph-url]

[![codecov][codecov-shield]][codecov-url]
[![Build][build-status-shield]][build-status-url]

rpiSystemMonitor is a tool to read vital system information of a RaspberryPi or other debian based distributions.

It is a web server and can be consumed by any application over the local network.

## Installation

run the `rpiSystemMonitor` binary

```bash
cd ./path/to/rpiSystemMonitor/
./rpiSystemMonitor

# alternatively you can run it in the background
./rpiSystemMonitor&
```

## API Routes

```bash
--localhost:3000/
  |-- cpu/ # TODO: CPU status as json
      |-- temp # CPU temp in C
      |-- freq # CPU freq in Hz
  |-- mem/
      |-- free # percent memory free
      |-- used # percent memory used
      |-- swp/
          |-- free # percent swap mem free
          |-- used # percent swap mem used
  |-- dsk/ # TODO
      |-- {ID}
          |-- free # TODO
          |-- used # TODO
```

## Usage

Open a browser on `localhost:3000/cpu/temp` to view cpu temps or `.../cpu/freq` to view cpu clock speed and so on.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

Distributed under the Apache License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[stars-shield]: https://img.shields.io/github/stars/RafaelZasas/rpiSystemMonitor.svg
[stars-url]: https://github.com/RafaelZasas/rpiSystemMonitor.git/stargazers
[forks-shield]: https://img.shields.io/github/forks/RafaelZasas/rpiSystemMonitor.svg
[forks-url]: https://github.com/RafaelZasas/rpiSystemMonitor.git/network/members
[sourcegraph-shield]: https://sourcegraph.com/github.com/rafaelzasas/rpiSystemMonitor/-/badge.svg
[sourcegraph-url]: https://sourcegraph.com/github.com/rafaelzasas/rpiSystemMonitor?badge
[codecov-shield]: https://codecov.io/gh/RafaelZasas/rpiSystemMonitor/branch/develop/graph/badge.svg?token=7WRZLTJJS6
[codecov-url]: https://codecov.io/gh/RafaelZasas/rpiSystemMonitor
[build-status-shield]: https://github.com/RafaelZasas/rpiSystemMonitor/actions/workflows/go.yml/badge.svg
[build-status-url]: https://github.com/RafaelZasas/rpiSystemMonitor/actions/workflows/go.yml