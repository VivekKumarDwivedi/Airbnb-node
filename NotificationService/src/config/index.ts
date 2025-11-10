//this file contain all the basic configuration logic to server work
import dotenv from 'dotenv';

type ServerConfig={
PORT:number,
REDIS_HOST?:string,
REDIS_PORT?:number,
MAIL_PASS?:string,
MAIL_USER?:string
};
 function loadEnv() {
  dotenv.config();
  console.log('Environment variables loaded successfully');
}
loadEnv();

export const serverConfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3000,
  REDIS_HOST: process.env.REDIS_HOST|| 'localhost',
  REDIS_PORT: process.env.REDIS_PORT ?Number(process.env.REDIS_PORT) :6379,
  MAIL_PASS: process.env.MAIL_PASS || '',
  MAIL_USER: process.env.MAIL_USER || ''
};
