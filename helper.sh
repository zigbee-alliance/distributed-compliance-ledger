#make localnet_stop
#
#make build
#make install
##make image
#
#sudo rm -rf localnet
#make localnet_init
##docker-compose up -d --no-deps --build
#make localnet_start

#########

make build
make install

sh genlocalconfig.sh
zbld start
