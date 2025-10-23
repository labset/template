import { createContext } from "react";

// eslint-disable-next-line @typescript-eslint/no-empty-object-type
export interface ConnectApi {}

export const ConnectApiContext = createContext<ConnectApi | undefined>(
  undefined,
);
