declare module "redlock" {
  import type { Redis } from "ioredis";

  export interface RedlockOptions {
    driftFactor?: number;
    retryCount?: number;
    retryDelay?: number;
    retryJitter?: number;
  }

  export interface Lock {
    unlock(): Promise<void>;
    extend(ttl: number): Promise<void>;
  }

  export default class Redlock {
    constructor(clients: Redis[], options?: RedlockOptions);

    acquire(resource: string | string[], ttl: number): Promise<Lock>;

    using<T>(
      resource: string | string[],
      ttl: number,
      routine: () => Promise<T>
    ): Promise<T>;
  }
}
