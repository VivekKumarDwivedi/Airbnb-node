import type  {Request, Response} from 'express';

export const pingHandler = async(req: Request, res: Response) => {
  res.send('pong');
};
