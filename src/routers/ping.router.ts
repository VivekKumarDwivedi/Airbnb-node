import express from "express";
import  { pingHandler } from "../controllers/ping.controller";
import type{ Request, Response, NextFunction } from "express";
const pingRouter = express.Router();


function middleware1(req: Request, res: Response, next: NextFunction) {
  // Middleware logic here
  console.log("Request received");
  next();
}

pingRouter.get('/ping', middleware1, pingHandler);

export default pingRouter;