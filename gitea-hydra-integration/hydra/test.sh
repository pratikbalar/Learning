 #!/bin/bash
 myls() {                                        
  docker-compose -f hydrastart.yml exec hydra  hydra clients create --endpoint http://127.0.0.1:4445 --id gitea-client --secret secret --grant-types authorization_code,refresh_token --response-types code,id_token --scope openid,offline --callbacks http://127.0.0.1:8000/user/oauth2/ORY_Hydra/callback --token-endpoint-auth-method client_secret_post
}
 myls1() {
  docker-compose -f hydrastart.yml exec hydra hydra token user --client-id gitea-client --client-secret secret --endpoint http://127.0.0.1:4444/ --port 5555 --scope openid,offline
}

myls
sleep 5
myls1
