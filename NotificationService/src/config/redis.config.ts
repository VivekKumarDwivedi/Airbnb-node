import Redis from 'ioredis';
import { serverConfig } from '.';

//Singleton pattern to connect to Redis
export function connectToRedis(){
    try{
        let connection:Redis;

        const redisConfig = {
         port : serverConfig.REDIS_PORT,
         host : serverConfig.REDIS_HOST,
         maxRetriesPerRequest: null,
        }
        return ()=>{
            if(!connection){
                connection = new Redis(redisConfig);
                return connection;
            }
            return connection;
        }
    }catch(error){
         console.log('Error connecting to Redis:',error);
         throw error;
    }
}

export const getRedisConnObject = connectToRedis();
