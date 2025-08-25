import express from 'express';
import  {serverConfig} from "./config/index";
import pingRouter from "./routers/ping.router";
const app = express();

/**
 * registering routes and there corresponding route without app server object
 */
app.use(pingRouter);

app.listen(serverConfig.PORT, () => {
  console.log(`Server is running on http://localhost:${serverConfig.PORT}`);
  console.log(`Press Ctrl+C to stop the server`);
});
