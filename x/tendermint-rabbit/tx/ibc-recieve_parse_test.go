package tx

import (
	"fmt"
	"testing"
	"time"
)

var old = []byte(`{"jsonrpc":"2.0","id":0,"result":{"query":"tm.event = 'Tx'","data":{"type":"tendermint\/event\/Tx","value":{"TxResult":{"height":"4536","index":0,"tx":"4RAoKBapCrwF+4NdCQoMaXJpc29rY2xpZW50EpEFCokECssCCgIIChIHaXJpc2h1Yhie9QwiDAjms\/z0BRCUzPSLAypICiCqS4VuVBAN8+BN5Lqzgo\/LJBsj3M6znKRYzlVfVh5I8hIkCAESIJXHZEgsI5IlF5WOOcNRzlANZHGseK1qJGHcZSWxonYmMiAVYRmDksmDRJokZV3D2A5FGesSwIOsqeFrbdUhsM\/T5EIgKmOToyFp+OTRb2Wm2d5invJ2sLF9hGkYZ+fijZmSCY9KICpjk6Mhafjk0W9lptneYp7ydrCxfYRpGGfn4o2ZkgmPUiAEgJG8fdwoP3e\/v5HXPETaWMPfipy8hnQF2Lfz2q2iL1ogVki0rGQ8hyeqLXvDYXInmOHxltxtTH9XHHtm7obQ50ViIG40C5z\/s3qYnKVE5rt4Cix4kB0\/szc4doURowYXr6AdchRqVCOdF8lEBUsz0LNUagXD6UhqshK4AQie9QwaSAog0a4qWyCDnW7ko29EQLwboJ4c0en1EYv0zrP5uCx1zzUSJAgBEiAXddJ+brqpfnHLy7Svp6fGQM1SsDmX5p06mZ6mjxAdMiJoCAISFGpUI50XyUQFSzPQs1RqBcPpSGqyGgwI67P89AUQor35mwMiQAgl8sG\/5vr8zuTG1P8X52NlsVk9goS6Ss5RG4Z3C4U10Q4S3vFkLAWQqOAzxXa1SsvqfHGPu4Dg4iENChnH6wQSggEKPwoUalQjnRfJRAVLM9CzVGoFw+lIarISJRYk3mQgscrYSXXZeOS2emnhF+0mTa\/fr5y+ILiEcsV55PeYdy8YZBI\/ChRqVCOdF8lEBUsz0LNUagXD6UhqshIlFiTeZCCxythJddl45LZ6aeEX7SZNr9+vnL4guIRyxXnk95h3LxhkGhSWZiSA7pZMamdRPIB3xZ7UtEKEWAqcCl1iymoKjwEKWbu9rx4KJAoadHJhbnNmZXIvb2tpcmlzY2hhbm5lbC9va3QSBjEwMDAwMBIUy4XFuHIrGMGcAuzAqpaFEgu9y0EaFJZmJIDulkxqZ1E8gHfFntS0QoRYKJ4rEAEaCHRyYW5zZmVyIg1va2lyaXNjaGFubmVsKgh0cmFuc2ZlcjINaXJpc29rY2hhbm5lbBLpCGrjE8MK4ggK0gMKBmlhdmw6dhIxNS9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9va2lyaXNjaGFubmVsL3BhY2tldHMvMRqUA5IDCo8DCioIDhA6GJ31DCog7LAcMSRqj9ShUcbIxRaJ90\/lDV0WqbcAetQgY5AIqXUKKggMEBoYnfUMKiDKzu6p1nneq1DFl2dzXHXgPaTgfeRPlqs4LaSYxgD+ywoqCAoQERid9QwqIJQhnmz5b0yzS4WQlFO28P5tlEZybg6j28+8IpLNYUapCioICBAJGJ31DCIgSpajwRztsE\/QM1yr82PaqtKtzsW9MIqkpwAWe+Qa5HsKKggGEAYYnfUMKiB4qPXi4WbtSfd5U0pLC8rww0vRsyihZZhfCiFdqG2zgAoqCAQQBBid9QwiIDJzpJzOyso9tU5a6sXAFKEEwa967JpFqtzrpEgemkupCioIAhACGJ31DCIgeJUzA7M1fISgbUaGj+8eCyR6YD8ZZZEqNdPyAZchonYaWQoxNS9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9va2lyaXNjaGFubmVsL3BhY2tldHMvMRIgqd3lfAN3Cm7J7xRCeYowcYvuYLg13nPAI2w1Mw8RY0YYnfUMCooFCgptdWx0aXN0b3JlEgNpYmMa9gT0BArxBAowCgRtYWluEigKJgid9QwSILGc8Ji0Ozx9VHmf7wlEqpUJn4Ldwhn6DICS5FKigr9BCjIKBnN1cHBseRIoCiYInfUMEiBWZozJVaLbtdh9jOUbNbSGtYhkyC6ssj3AQeh4I\/M4HAoRCgd1cGdyYWRlEgYKBAid9QwKMwoHc3Rha2luZxIoCiYInfUMEiBLGE+TUoGvI\/Iu2Ug3SD4VsDHGS64LZ2arQNcgOY3FxAoyCgZwYXJhbXMSKAomCJ31DBIgmWIBsU6ErA3514kv0lp8lmq5zYrFQzV\/NHn9CHdwD\/4KMAoEbWludBIoCiYInfUMEiCWSGhLOvR5O2fTcdeTNz3O9YHFiGamKrSCqRWbxkGo6wowCgRiYW5rEigKJgid9QwSIGr53GxpVEgyv3p1gSMxQcR7qHaXs5MpQGV2vADZ3i6dCjQKCHNsYXNoaW5nEigKJgid9QwSID7NiYuG2RWwQZKJmVmrrYn3TB4+0MsPhlCgiqiU\/t9HCi8KA2dvdhIoCiYInfUMEiBgzRN8GWLsrGFjidaANIM\/KSFQnC0oWqX1FTmXzpaKdAovCgNhY2MSKAomCJ31DBIgdrM94QnLsWwXAtFkyMfRAb069jgsGS3hJg3zZ4EAie0KOAoMZGlzdHJpYnV0aW9uEigKJgid9QwSIOVRnxwFYbRwpdmpcl36C+\/EWxCFuVCLIsXFYsYukE\/WCi8KA2liYxIoCiYInfUMEiAGA+MnCZ8useWISKzN+OzeKfH+bl+aoEr04T5cpnWFugoSCghldmlkZW5jZRIGCgQInfUMChIKCHRyYW5zZmVyEgYKBAid9QwYnvUMIhSWZiSA7pZMamdRPIB3xZ7UtEKEWBIRCgsKA29rdBIENTAwMBDAmgwaagom61rphyEDKIb8W0m16DjKZTsWDNer55xNUnnbt33NXvp2\/N3kr38SQJ0kWlsH0mzj1oUxhNAIUe5SRu+RbutCTksw5Xzr2WvfOmXTZz7sR0jtLV3wXv7QOsHmhVi0kg3lL28N5J0y5JY=","result":{"log":"[{\"msg_index\":0,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"update_client\"}]}]},{\"msg_index\":1,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"ics20\/transfer\"},{\"key\":\"sender\",\"value\":\"cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592\"},{\"key\":\"module\",\"value\":\"ibc_client\"},{\"key\":\"sender\",\"value\":\"cosmos1ekg5kmfx5qeaxyu3fve2us0y79g6hm7t6tphm5\"},{\"key\":\"module\",\"value\":\"ibc_transfer\"},{\"key\":\"sender\",\"value\":\"cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta\"},{\"key\":\"amount\",\"value\":\"5000okt\"},{\"key\":\"recipient\",\"value\":\"cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592\"},{\"key\":\"amount\",\"value\":\"100000okt\"}]},{\"type\":\"update_client\",\"attributes\":[{\"key\":\"client_id\",\"value\":\"irisokclient\"}]}]}]","gas_wanted":"200000","gas_used":"121812","events":[{"type":"message","attributes":[{"key":"YWN0aW9u","value":"dXBkYXRlX2NsaWVudA=="}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"aWNzMjAvdHJhbnNmZXI="}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y29zbW9zMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsc2VycXRh"},{"key":"YW1vdW50","value":"NTAwMG9rdA=="}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y29zbW9zMWplbnpmcTh3amV4eDVlNjM4anE4MDN2NzZqNnk5cHpjeWQ4NTky"}]},{"type":"update_client","attributes":[{"key":"Y2xpZW50X2lk","value":"aXJpc29rY2xpZW50"}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"aWJjX2NsaWVudA=="}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y29zbW9zMWplbnpmcTh3amV4eDVlNjM4anE4MDN2NzZqNnk5cHpjeWQ4NTky"},{"key":"YW1vdW50","value":"MTAwMDAwb2t0"}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y29zbW9zMWVrZzVrbWZ4NXFlYXh5dTNmdmUydXMweTc5ZzZobTd0NnRwaG01"}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"aWJjX3RyYW5zZmVy"},{"key":"c2VuZGVy","value":"Y29zbW9zMWplbnpmcTh3amV4eDVlNjM4anE4MDN2NzZqNnk5cHpjeWQ4NTky"}]}]}}}},"events":{"message.action":["update_client","ics20\/transfer"],"transfer.recipient":["cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta","cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592"],"transfer.amount":["5000okt","100000okt"],"message.sender":["cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592","cosmos1ekg5kmfx5qeaxyu3fve2us0y79g6hm7t6tphm5","cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592"],"update_client.client_id":["irisokclient"],"tx.height":["4536"],"message.module":["ibc_client","ibc_transfer"],"tm.event":["Tx"],"tx.hash":["D42DEE12016AC872BD3407BCF6DC045468E01B32EF40E0C7BB821E0358AE85A7"]}}}`)
var new = []byte(`{"jsonrpc":"2.0","id":0,"result":{"query":"tm.event = 'Tx'","data":{"type":"tendermint\/event\/Tx","value":{"TxResult":{"height":"4536","index":0,"tx":"4RAoKBapCrwF+4NdCQoMaXJpc29rY2xpZW50EpEFCokECssCCgIIChIHaXJpc2h1Yhie9QwiDAjms\/z0BRCUzPSLAypICiCqS4VuVBAN8+BN5Lqzgo\/LJBsj3M6znKRYzlVfVh5I8hIkCAESIJXHZEgsI5IlF5WOOcNRzlANZHGseK1qJGHcZSWxonYmMiAVYRmDksmDRJokZV3D2A5FGesSwIOsqeFrbdUhsM\/T5EIgKmOToyFp+OTRb2Wm2d5invJ2sLF9hGkYZ+fijZmSCY9KICpjk6Mhafjk0W9lptneYp7ydrCxfYRpGGfn4o2ZkgmPUiAEgJG8fdwoP3e\/v5HXPETaWMPfipy8hnQF2Lfz2q2iL1ogVki0rGQ8hyeqLXvDYXInmOHxltxtTH9XHHtm7obQ50ViIG40C5z\/s3qYnKVE5rt4Cix4kB0\/szc4doURowYXr6AdchRqVCOdF8lEBUsz0LNUagXD6UhqshK4AQie9QwaSAog0a4qWyCDnW7ko29EQLwboJ4c0en1EYv0zrP5uCx1zzUSJAgBEiAXddJ+brqpfnHLy7Svp6fGQM1SsDmX5p06mZ6mjxAdMiJoCAISFGpUI50XyUQFSzPQs1RqBcPpSGqyGgwI67P89AUQor35mwMiQAgl8sG\/5vr8zuTG1P8X52NlsVk9goS6Ss5RG4Z3C4U10Q4S3vFkLAWQqOAzxXa1SsvqfHGPu4Dg4iENChnH6wQSggEKPwoUalQjnRfJRAVLM9CzVGoFw+lIarISJRYk3mQgscrYSXXZeOS2emnhF+0mTa\/fr5y+ILiEcsV55PeYdy8YZBI\/ChRqVCOdF8lEBUsz0LNUagXD6UhqshIlFiTeZCCxythJddl45LZ6aeEX7SZNr9+vnL4guIRyxXnk95h3LxhkGhSWZiSA7pZMamdRPIB3xZ7UtEKEWAqcCl1iymoKjwEKWbu9rx4KJAoadHJhbnNmZXIvb2tpcmlzY2hhbm5lbC9va3QSBjEwMDAwMBIUy4XFuHIrGMGcAuzAqpaFEgu9y0EaFJZmJIDulkxqZ1E8gHfFntS0QoRYKJ4rEAEaCHRyYW5zZmVyIg1va2lyaXNjaGFubmVsKgh0cmFuc2ZlcjINaXJpc29rY2hhbm5lbBLpCGrjE8MK4ggK0gMKBmlhdmw6dhIxNS9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9va2lyaXNjaGFubmVsL3BhY2tldHMvMRqUA5IDCo8DCioIDhA6GJ31DCog7LAcMSRqj9ShUcbIxRaJ90\/lDV0WqbcAetQgY5AIqXUKKggMEBoYnfUMKiDKzu6p1nneq1DFl2dzXHXgPaTgfeRPlqs4LaSYxgD+ywoqCAoQERid9QwqIJQhnmz5b0yzS4WQlFO28P5tlEZybg6j28+8IpLNYUapCioICBAJGJ31DCIgSpajwRztsE\/QM1yr82PaqtKtzsW9MIqkpwAWe+Qa5HsKKggGEAYYnfUMKiB4qPXi4WbtSfd5U0pLC8rww0vRsyihZZhfCiFdqG2zgAoqCAQQBBid9QwiIDJzpJzOyso9tU5a6sXAFKEEwa967JpFqtzrpEgemkupCioIAhACGJ31DCIgeJUzA7M1fISgbUaGj+8eCyR6YD8ZZZEqNdPyAZchonYaWQoxNS9wb3J0cy90cmFuc2Zlci9jaGFubmVscy9va2lyaXNjaGFubmVsL3BhY2tldHMvMRIgqd3lfAN3Cm7J7xRCeYowcYvuYLg13nPAI2w1Mw8RY0YYnfUMCooFCgptdWx0aXN0b3JlEgNpYmMa9gT0BArxBAowCgRtYWluEigKJgid9QwSILGc8Ji0Ozx9VHmf7wlEqpUJn4Ldwhn6DICS5FKigr9BCjIKBnN1cHBseRIoCiYInfUMEiBWZozJVaLbtdh9jOUbNbSGtYhkyC6ssj3AQeh4I\/M4HAoRCgd1cGdyYWRlEgYKBAid9QwKMwoHc3Rha2luZxIoCiYInfUMEiBLGE+TUoGvI\/Iu2Ug3SD4VsDHGS64LZ2arQNcgOY3FxAoyCgZwYXJhbXMSKAomCJ31DBIgmWIBsU6ErA3514kv0lp8lmq5zYrFQzV\/NHn9CHdwD\/4KMAoEbWludBIoCiYInfUMEiCWSGhLOvR5O2fTcdeTNz3O9YHFiGamKrSCqRWbxkGo6wowCgRiYW5rEigKJgid9QwSIGr53GxpVEgyv3p1gSMxQcR7qHaXs5MpQGV2vADZ3i6dCjQKCHNsYXNoaW5nEigKJgid9QwSID7NiYuG2RWwQZKJmVmrrYn3TB4+0MsPhlCgiqiU\/t9HCi8KA2dvdhIoCiYInfUMEiBgzRN8GWLsrGFjidaANIM\/KSFQnC0oWqX1FTmXzpaKdAovCgNhY2MSKAomCJ31DBIgdrM94QnLsWwXAtFkyMfRAb069jgsGS3hJg3zZ4EAie0KOAoMZGlzdHJpYnV0aW9uEigKJgid9QwSIOVRnxwFYbRwpdmpcl36C+\/EWxCFuVCLIsXFYsYukE\/WCi8KA2liYxIoCiYInfUMEiAGA+MnCZ8useWISKzN+OzeKfH+bl+aoEr04T5cpnWFugoSCghldmlkZW5jZRIGCgQInfUMChIKCHRyYW5zZmVyEgYKBAid9QwYnvUMIhSWZiSA7pZMamdRPIB3xZ7UtEKEWBIRCgsKA29rdBIENTAwMBDAmgwaagom61rphyEDKIb8W0m16DjKZTsWDNer55xNUnnbt33NXvp2\/N3kr38SQJ0kWlsH0mzj1oUxhNAIUe5SRu+RbutCTksw5Xzr2WvfOmXTZz7sR0jtLV3wXv7QOsHmhVi0kg3lL28N5J0y5JY=","result":{"log":"[{\"msg_index\":0,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"update_client\"}]}]},{\"msg_index\":1,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"ics20\/transfer\"},{\"key\":\"sender\",\"value\":\"cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592\"},{\"key\":\"module\",\"value\":\"ibc_client\"},{\"key\":\"sender\",\"value\":\"cosmos1ekg5kmfx5qeaxyu3fve2us0y79g6hm7t6tphm5\"},{\"key\":\"module\",\"value\":\"ibc_transfer\"},{\"key\":\"sender\",\"value\":\"cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta\"},{\"key\":\"amount\",\"value\":\"5000okt\"},{\"key\":\"recipient\",\"value\":\"cosmos1jenzfq8wjexx5e638jq803v76j6y9pzcyd8592\"},{\"key\":\"amount\",\"value\":\"100000okt\"}]},{\"type\":\"update_client\",\"attributes\":[{\"key\":\"client_id\",\"value\":\"irisokclient\"}]}]}]","gas_wanted":"200000","gas_used":"121812","events":[{"type":"message","attributes":[{"key":"YWN0aW9u","value":"dXBkYXRlX2NsaWVudA=="}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"aWNzMjAvdHJhbnNmZXI="}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y29zbW9zMTd4cGZ2YWttMmFtZzk2MnlsczZmODR6M2tlbGw4YzVsc2VycXRh"},{"key":"YW1vdW50","value":"NTAwMG9rdA=="}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y29zbW9zMWplbnpmcTh3amV4eDVlNjM4anE4MDN2NzZqNnk5cHpjeWQ4NTky"}]},{"type":"update_client","attributes":[{"key":"Y2xpZW50X2lk","value":"aXJpc29rY2xpZW50"}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"aWJjX2NsaWVudA=="}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"Y29zbW9zMWplbnpmcTh3amV4eDVlNjM4anE4MDN2NzZqNnk5cHpjeWQ4NTky"},{"key":"YW1vdW50","value":"MTAwMDAwb2t0"}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"Y29zbW9zMWVrZzVrbWZ4NXFlYXh5dTNmdmUydXMweTc5ZzZobTd0NnRwaG01"}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"aWJjX3RyYW5zZmVy"},{"key":"c2VuZGVy","value":"Y29zbW9zMWplbnpmcTh3amV4eDVlNjM4anE4MDN2NzZqNnk5cHpjeWQ4NTky"}]}]}}}},"events":{"transfer.recipient":["cosmos1u3gzlq85frukjfwsvgnn7rs83xvfhrkhj2uj7l"],"transfer.amount":["10000n0token"],"recv_packet.packet_dst_port":["transfer"],"fungible_token_packet.receiver":["cosmos1u3gzlq85frukjfwsvgnn7rs83xvfhrkhj2uj7l"],"tm.event":["Tx"],"tx.height":["90"],"message.sender":["cosmos1qmvw4k4rgu53066yjaz03m83uzvexhytsjq4er"],"recv_packet.packet_timeout":["1087"],"fungible_token_packet.module":["ibc_transfer"],"tx.hash":["16BA8625F42D1673453D0EEF6B8155FFD9B1439668281F1B1FDA38647AF81D71"],"message.action":["update_client","ics04\/opaque"],"recv_packet.packet_sequence":["1"],"recv_packet.packet_src_port":["transfer"],"recv_packet.packet_src_channel":["ibczeroxfer"],"fungible_token_packet.value":["10000transfer\/ibczeroxfer\/n0token"],"recv_packet.packet_data":["{\"error\":\"\",\"success\":true}"],"recv_packet.packet_dst_channel":["ibconexfer"]}}}`)

func TestIbcRecieveParse(t *testing.T) {
	oldTx, err := ParseTx(old)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(oldTx.Normalize(time.Now(), "test", 0).Marshal()))

	newTx, err := ParseTx(new)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(newTx.Normalize(time.Now(), "test", 0).Marshal()))
}