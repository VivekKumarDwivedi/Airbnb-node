import IORedis from 'ioredis';
import RedLock from 'redlock';
import { serverConfig } from '.';

export const redisClient = new IORedis(serverConfig.REDIS_SERVER_URL);

export const redlock = new RedLock([redisClient],{
    driftFactor:0.01,
    retryCount:10,
    retryDelay:200,
    retryJitter:200
});

