import express from 'express';
import  {serverConfig} from "./config/index";
import v1Router from "./routers/v1/index.router";
import v2Router from "./routers/v2/index.router";
import { genericErrorHandler } from './middlewares/error.middleware';
import logger from './config/logger.config';
import {  attachCorrelationIdMiddleware } from './middlewares/correlation.middleware';
import { setupMailerWorker } from './processors/email.processor';
import { addEmailToQueue } from './producers/email.producer';
const app = express();

app.use(express.json());
/**
 * registering routes and there corresponding route without app server object
 */
app.use(attachCorrelationIdMiddleware);


app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router); // just for testing purpose we are using v1 router for v2 also


/**
 * registering error middleware
 */
app.use(genericErrorHandler);
app.listen(serverConfig.PORT, async() => {
  logger.info(`Server is running on http://localhost:${serverConfig.PORT}`);
  logger.info(`Press Ctrl+C to stop the server`);
  setupMailerWorker();
  logger.info(`Mailer worker setup completed.`);
   
  addEmailToQueue({
    to:"vivekdwivedi2023@gmail.com",
    subject:"Test Email",
    templateId:"welcome",
    params:{
      name: "Vivek Dwivedi",
      appName: "Booking App",
    }
  })

});
