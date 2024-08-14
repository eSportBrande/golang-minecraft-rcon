# MineCraft RCON using GoLang.
Building using the https://github.com/FragLand/minestat/ library for Go.

Used for getting players and version data in a clean output for use in other scripts in our usecase for a PowerShell script and also for InfluxDB.

Check the releases for the latest version for your platform.

## Usage
### Get Players only:
```bash
chmod +x ./minecraft-rcon
./minecraft-rcon -serverIPorDNS 127.0.0.1 -playersOnly=true
```

### Get Version only:
```bash
chmod +x ./minecraft-rcon
./minecraft-rcon -serverIPorDNS 127.0.0.1 -versionOnly=true
```

### Get Players and Version:
```bash
chmod +x ./minecraft-rcon
./minecraft-rcon -serverIPorDNS 127.0.0.1
```
