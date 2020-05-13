#!/bin/bash
watcher --tmRPC "tcp://3.21.169.1:26657" --rabbitMQ "$RABBITMQ" --zone zilotchain &
watcher --tmRPC "tcp://ibc.westaking.io:26657" --rabbitMQ "$RABBITMQ" --zone westaking &
watcher --tmRPC "tcp://chain.exchange-fees.com:26657" --rabbitMQ "$RABBITMQ" --zone vostok-1 &
watcher --tmRPC "tcp://173.249.12.108:26657" --rabbitMQ "$RABBITMQ" --zone vipnamai &
watcher --tmRPC "tcp://ibc.vgng.io:26657" --rabbitMQ "$RABBITMQ" --zone vgng-1 &
watcher --tmRPC "tcp://144.202.100.245:26657" --rabbitMQ "$RABBITMQ" --zone vbstreetz &
watcher --tmRPC "tcp://goz.umbrellavalidator.com:26657" --rabbitMQ "$RABBITMQ" --zone umbrella &
watcher --tmRPC "tcp://3.22.166.56:26657" --rabbitMQ "$RABBITMQ" --zone timetowinchain &
watcher --tmRPC "tcp://thetechtrap.com:26657" --rabbitMQ "$RABBITMQ" --zone thetechtrap-goz &
watcher --tmRPC "tcp://supernovachain.cosmostation.io:26657" --rabbitMQ "$RABBITMQ" --zone supernovachain &
watcher --tmRPC "tcp:stratus.mycryptobets.com:26657" --rabbitMQ "$RABBITMQ" --zone stardust-1111 &
watcher --tmRPC "tcp://goz.starcluster.tech:26657" --rabbitMQ "$RABBITMQ" --zone starcluster-1337 &
watcher --tmRPC "tcp://161.35.45.178:26657" --rabbitMQ "$RABBITMQ" --zone stakin &
watcher --tmRPC "tcp://goz2.stake.zone:26657" --rabbitMQ "$RABBITMQ" --zone szchain &
watcher --tmRPC "tcp://stakewolf.com:26657" --rabbitMQ "$RABBITMQ" --zone stakewolf &
watcher --tmRPC "tcp://35.198.125.128:26657" --rabbitMQ "$RABBITMQ" --zone stakematic &
watcher --tmRPC "tcp://goz.cosmos.fish:26657" --rabbitMQ "$RABBITMQ" --zone jellyfish &
watcher --tmRPC "tcp://ibc.staked.cloud:26657" --rabbitMQ "$RABBITMQ" --zone staked-ibc &
watcher --tmRPC "tcp://goz.stakebird.com:26657" --rabbitMQ "$RABBITMQ" --zone stakebird-1 &
watcher --tmRPC "tcp://ibc.stake.sh:26657" --rabbitMQ "$RABBITMQ" --zone nuube-goz &
watcher --tmRPC "tcp://goz.stakedao.org:26657" --rabbitMQ "$RABBITMQ" --zone stake-capital &
watcher --tmRPC "tcp://goz.source.network:26657" --rabbitMQ "$RABBITMQ" --zone source-xgoz &
watcher --tmRPC "tcp://ananas.alpe1.net:26657" --rabbitMQ "$RABBITMQ" --zone snakey &
watcher --tmRPC "tcp://80.64.211.64:26657" --rabbitMQ "$RABBITMQ" --zone simplystaking &
watcher --tmRPC "tcp://setan.ml:26657" --rabbitMQ "$RABBITMQ" --zone setanchain &
watcher --tmRPC "tcp://one.goz.sentinel.co:26657" --rabbitMQ "$RABBITMQ" --zone sentinel-goz &
watcher --tmRPC "tcp://tnet-csg.c9ret.xyz:26657" --rabbitMQ "$RABBITMQ" --zone retz80chain &
watcher --tmRPC "tcp://regengoz.vaasl.io:26657" --rabbitMQ "$RABBITMQ" --zone regengoz &
watcher --tmRPC "tcp://node.pylons.tech:26657" --rabbitMQ "$RABBITMQ" --zone pylonschain &
watcher --tmRPC "tcp://35.230.42.221:26657" --rabbitMQ "$RABBITMQ" --zone pupu &
watcher --tmRPC "tcp://3.134.115.80:26657" --rabbitMQ "$RABBITMQ" --zone Protofire.io &
watcher --tmRPC "tcp://ibc.j96.me:26657" --rabbitMQ "$RABBITMQ" --zone plex &
watcher --tmRPC "tcp://ibc.ping.pub:26657" --rabbitMQ "$RABBITMQ" --zone ping-ibc &
watcher --tmRPC "tcp://goz.cyphercore.io:26657" --rabbitMQ "$RABBITMQ" --zone petomhub &
watcher --tmRPC "tcp://relayer.persistence.one:26657" --rabbitMQ "$RABBITMQ" --zone persistence &
watcher --tmRPC "tcp://p2p-org-1.goz.p2p.org:26657" --rabbitMQ "$RABBITMQ" --zone p2p-org-1 &
watcher --tmRPC "tcp://goz.ozonechain.xyz:26657" --rabbitMQ "$RABBITMQ" --zone ozone &
watcher --tmRPC "tcp://goz.kysenpool.io:26657" --rabbitMQ "$RABBITMQ" --zone outpost &
watcher --tmRPC "tcp://3.112.29.150:26657" --rabbitMQ "$RABBITMQ" --zone okchain &
watcher --tmRPC "tcp://144.76.118.133:26657" --rabbitMQ "$RABBITMQ" --zone nodeasy &
watcher --tmRPC "tcp://193.30.121.61:26657" --rabbitMQ "$RABBITMQ" --zone Node123 &
watcher --tmRPC "tcp://goz.nibiru.network:26657" --rabbitMQ "$RABBITMQ" --zone nibiru-ibc &
watcher --tmRPC "tcp://muzamint.com:26657/" --rabbitMQ "$RABBITMQ" --zone muzamint &
watcher --tmRPC "tcp://goz.modulus.network:26657" --rabbitMQ "$RABBITMQ" --zone modulus-goz-1 &
watcher --tmRPC "tcp://mmmh.sytes.net:26657" --rabbitMQ "$RABBITMQ" --zone mmmh-lazy &
watcher --tmRPC "tcp://45.77.91.232:26657" --rabbitMQ "$RABBITMQ" --zone mintonium &
watcher --tmRPC "tcp://melea.xyz:26657" --rabbitMQ "$RABBITMQ" --zone melea-1111 &
watcher --tmRPC "tcp://validating.for.co.ke:26657" --rabbitMQ "$RABBITMQ" --zone meeseeks &
watcher --tmRPC "tcp://goz.labeleet.com:26657" --rabbitMQ "$RABBITMQ" --zone kugs-030 &
watcher --tmRPC "tcp://goz.konstellation.tech:26657" --rabbitMQ "$RABBITMQ" --zone konstellation &
watcher --tmRPC "tcp://goz.kiraex.com:10001" --rabbitMQ "$RABBITMQ" --zone kira-1 &
watcher --tmRPC "tcp://213.32.70.133:26657" --rabbitMQ "$RABBITMQ" --zone jtbchain &
watcher --tmRPC "tcp://54.211.26.151:26657" --rabbitMQ "$RABBITMQ" --zone js &
watcher --tmRPC "tcp://joon-chain-goz.cosmostation.io:26657" --rabbitMQ "$RABBITMQ" --zone joon-chain-goz &
watcher --tmRPC "tcp://49.12.106.6:26657" --rabbitMQ "$RABBITMQ" --zone isillienchain &
watcher --tmRPC "tcp://interstation.cosmostation.io:26657" --rabbitMQ "$RABBITMQ" --zone interstation &
watcher --tmRPC "tcp://18.178.211.15:26657" --rabbitMQ "$RABBITMQ" --zone hongo-3 &
watcher --tmRPC "tcp://goz.hashquark.io:26657" --rabbitMQ "$RABBITMQ" --zone hashquarkchain &
watcher --tmRPC "tcp://15.165.120.204:26657" --rabbitMQ "$RABBITMQ" --zone hada &
watcher --tmRPC "tcp://goz.gunray.xyz:26657" --rabbitMQ "$RABBITMQ" --zone gunchain &
watcher --tmRPC "tcp://176.9.8.110:26657" --rabbitMQ "$RABBITMQ" --zone grbx-route &
watcher --tmRPC "tcp://95.216.216.117:26657" --rabbitMQ "$RABBITMQ" --zone genesis-lab &
watcher --tmRPC "tcp://goz.jptpool.com:26657" --rabbitMQ "$RABBITMQ" --zone gemstone &
watcher --tmRPC "tcp://ibc.freeflix.media:26657" --rabbitMQ "$RABBITMQ" --zone freeflix-media-hub &
watcher --tmRPC "tcp://18.217.240.174:26657" --rabbitMQ "$RABBITMQ" --zone finalbattlechain &
watcher --tmRPC "tcp:/goz-ibc.figment.network:26657" --rabbitMQ "$RABBITMQ" --zone figment &
watcher --tmRPC "tcp://fetch-goz.fetch.ai:26657" --rabbitMQ "$RABBITMQ" --zone fetchBeacon &
watcher --tmRPC "tcp://15.236.69.21:26657" --rabbitMQ "$RABBITMQ" --zone ublochain &
watcher --tmRPC "tcp://35.209.174.13:26657" --rabbitMQ "$RABBITMQ" --zone stateset &
watcher --tmRPC "tcp://3.22.194.241:26657" --rabbitMQ "$RABBITMQ" --zone EVM.Protofire.io &
watcher --tmRPC "tcp://goz.everstake.one:26657" --rabbitMQ "$RABBITMQ" --zone everstakechain &
watcher --tmRPC "tcp://52.231.28.219:26657" --rabbitMQ "$RABBITMQ" --zone dunhillchain &
watcher --tmRPC "tcp://goz.dos.network:26657" --rabbitMQ "$RABBITMQ" --zone dos-ibc &
watcher --tmRPC "tcp://3.21.156.79:26657" --rabbitMQ "$RABBITMQ" --zone disraptorchain &
watcher --tmRPC "http://goz.desmos.network:80" --rabbitMQ "$RABBITMQ" --zone morpheus-goz-1a &
watcher --tmRPC "tcp://ibc.defending.network:26657" --rabbitMQ "$RABBITMQ" --zone defending-network &
watcher --tmRPC "tcp://goz.dev.datachain.jp:26657" --rabbitMQ "$RABBITMQ" --zone dcz &
watcher --tmRPC "tcp://testnet.dawns.world:26657" --rabbitMQ "$RABBITMQ" --zone dawnsworld &
watcher --tmRPC "http://cyberdevs.cyberd.ai:26657" --rabbitMQ "$RABBITMQ" --zone cyberdevs &
watcher --tmRPC "tcp://47.240.29.196:26657" --rabbitMQ "$RABBITMQ" --zone crazyzoo &
watcher --tmRPC "tcp://ibc.cosmoon.org:26657" --rabbitMQ "$RABBITMQ" --zone cosmoon-testnet &
watcher --tmRPC "tcp://supernova.commonwealth.im:26657" --rabbitMQ "$RABBITMQ" --zone supernova &
watcher --tmRPC "tcp://ibc.cosmiccompass.io:26657" --rabbitMQ "$RABBITMQ" --zone coco-post-chain &
watcher --tmRPC "tcp://fedzone.chorus.one:26657" --rabbitMQ "$RABBITMQ" --zone fedzone-1 &
watcher --tmRPC "tcp://176.9.238.157:26657" --rabbitMQ "$RABBITMQ" --zone chainlayer &
watcher --tmRPC "tcp://goz.chainflow.io:26657" --rabbitMQ "$RABBITMQ" --zone Chainflow &
watcher --tmRPC "http://goz.chainapsis.com:80" --rabbitMQ "$RABBITMQ" --zone chainapsis-1a &
watcher --tmRPC "tcp://wasntus.goz.certus.one:26657" --rabbitMQ "$RABBITMQ" --zone it-wasnt-us &
watcher --tmRPC "tcp://3.130.208.130:26657" --rabbitMQ "$RABBITMQ" --zone cat &
watcher --tmRPC "tcp://capychain.com:26657" --rabbitMQ "$RABBITMQ" --zone capychain &
watcher --tmRPC "tcp://178.128.186.74:26657" --rabbitMQ "$RABBITMQ" --zone byzantine-gos &
watcher --tmRPC "tcp://95.216.198.111:26657" --rabbitMQ "$RABBITMQ" --zone burek &
watcher --tmRPC "tcp://149.248.61.193:26657" --rabbitMQ "$RABBITMQ" --zone brankochain &
watcher --tmRPC "tcp://ibc.blockscape.network:26657" --rabbitMQ "$RABBITMQ" --zone NoChainNoGain-1000 &
watcher --tmRPC "tcp://rpc.blockngine.io:26657" --rabbitMQ "$RABBITMQ" --zone blockngine-ibc &
watcher --tmRPC "tcp://goz-1.bitsong.network:26657" --rabbitMQ "$RABBITMQ" --zone bitsong-goz-1 &
watcher --tmRPC "tcp://ibc.bharvest.io:26657" --rabbitMQ "$RABBITMQ" --zone B-Harvest &
watcher --tmRPC "tcp://blogchain.xyz:26657" --rabbitMQ "$RABBITMQ" --zone BlogChain &
watcher --tmRPC "tcp://52.247.127.40:26657" --rabbitMQ "$RABBITMQ" --zone aynchain &
watcher --tmRPC "tcp://goz.audit.one:26657" --rabbitMQ "$RABBITMQ" --zone audit.one &
watcher --tmRPC "tcp://bombers.dokia.cloud:26657" --rabbitMQ "$RABBITMQ" --zone atomicbombers &
watcher --tmRPC "tcp://bombers.dokia.cloud:26657" --rabbitMQ "$RABBITMQ" --zone atomic-bombers &
watcher --tmRPC "tcp://goz.armyids.com:26657" --rabbitMQ "$RABBITMQ" --zone armyids &
watcher --tmRPC "tcp://116.203.208.175:26657" --rabbitMQ "$RABBITMQ" --zone ape_smash &
watcher --tmRPC "tcp://116.203.245.68:26657" --rabbitMQ "$RABBITMQ" --zone anon-chain &
watcher --tmRPC "tcp://goz.aneka.io:26657" --rabbitMQ "$RABBITMQ" --zone aneka &
watcher --tmRPC "tcp://goz.cosmos.alphavirtual.com:26657" --rabbitMQ "$RABBITMQ" --zone avnet &
watcher --tmRPC "tcp://159.89.183.144:26657" --rabbitMQ "$RABBITMQ" --zone akashian &
watcher --tmRPC "tcp://138.197.157.152:26657" --rabbitMQ "$RABBITMQ" --zone aib-goz-1 &
watcher --tmRPC "tcp://178.128.254.143:26657" --rabbitMQ "$RABBITMQ" --zone agoric-goz-1.0.0 &
watcher --tmRPC "tcp://goz.wetez.io:26657" --rabbitMQ "$RABBITMQ" --zone Wetez &
watcher --tmRPC "tcp://47.245.35.172:26657" --rabbitMQ "$RABBITMQ" --zone tom &
watcher --tmRPC "tcp://34.66.86.162:26657" --rabbitMQ "$RABBITMQ" --zone dking &
watcher --tmRPC "tcp://13.250.207.24:26657" --rabbitMQ "$RABBITMQ" --zone TaidiHub &
watcher --tmRPC "tcp://ibc.staking.fund:26657" --rabbitMQ "$RABBITMQ" --zone stakingfund &
watcher --tmRPC "tcp://5.181.51.80:26657" --rabbitMQ "$RABBITMQ" --zone stakesstonechain &
watcher --tmRPC "tcp://18.136.225.71:26657" --rabbitMQ "$RABBITMQ" --zone SNZPoolHub &
watcher --tmRPC "tcp://rvc.novy.pw:26657" --rabbitMQ "$RABBITMQ" --zone rvc-1 &
watcher --tmRPC "tcp://goz.newroad.network:26657" --rabbitMQ "$RABBITMQ" --zone newroadchain &
watcher --tmRPC "tcp://35.193.176.142:26657" --rabbitMQ "$RABBITMQ" --zone universe &
watcher --tmRPC "tcp://molecule.adri.co:26657" --rabbitMQ "$RABBITMQ" --zone moleculechain &
watcher --tmRPC "tcp://45.79.207.112:26657" --rabbitMQ "$RABBITMQ" --zone microtick-ibc &
watcher --tmRPC "tcp://35.236.168.104:26657" --rabbitMQ "$RABBITMQ" --zone irishub-goz &
watcher --tmRPC "tcp://35.222.132.154:26657" --rabbitMQ "$RABBITMQ" --zone seen &
watcher --tmRPC "tcp://val1.goz.enigma.co:26657" --rabbitMQ "$RABBITMQ" --zone enigma-goz &
watcher --tmRPC "tcp://dropschain.com:26657" --rabbitMQ "$RABBITMQ" --zone dropschain &
watcher --tmRPC "tcp://goz.val.network:26657" --rabbitMQ "$RABBITMQ" --zone Compass &
watcher --tmRPC "tcp://achain.nodeateam.com:26657" --rabbitMQ "$RABBITMQ" --zone achain &
watcher --tmRPC "ws://35.233.155.199:26657/websocket" --rabbitMQ "$RABBITMQ" --zone gameofzoneshub-1a &
wait