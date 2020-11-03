```     
docker-compose -f hydrastart.yml \
    -f hydrastart-postgres.yml \
    up --build 
    
 
 docker-compose -f hydrastart.yml exec hydra \
    hydra clients create \
    --endpoint http://127.0.0.1:4445/ \
    --id my-client \
    --secret secret \
    -g client_credentials 
    
 sudo docker-compose -f hydrastart.yml exec hydra \
    hydra token client \
    --endpoint http://127.0.0.1:4444/ \
    --client-id my-client \
    --client-secret secret  ```

