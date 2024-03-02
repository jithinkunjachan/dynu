# dynu
This go binary helps to update the ip address in https://www.dynu.com/ every 50 seconds


## Installation

### Download the binary
`wget https://github.com/jithinkunjachan/dynu/releases/download/1.0.2/dynu-arm64`

### Make binary executable
`chmod +x dynu-arm64`

### Create a shell script
`touch dynu.sh`

#### Enter following in `dynu.sh`
```
export DYNU_USERNAME=xxxx
export DYNU_PASSWORD=passwordxxx`
/home/rj/dynu-arm64
 ```

### Update the crontab
`sudo crontab -e`

#### Enter following in crontab
```
* * * * * /home/rj/dynu.sh
```
