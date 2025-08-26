import express from 'express';
import  {serverConfig} from "./config/index";
import v1Router from "./routers/v1/index.router";
import v2Router from './routers/v2/index.router';
const app = express();

/**
 * registering routes and there corresponding route without app server object
 */
app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router); // just for testing purpose we are using v1 router for v2 also

app.listen(serverConfig.PORT, () => {
  console.log(`Server is running on http://localhost:${serverConfig.PORT}`);
  console.log(`Press Ctrl+C to stop the server`);
});
