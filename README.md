# rpiSystemMonitor

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

```
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
// TODO
