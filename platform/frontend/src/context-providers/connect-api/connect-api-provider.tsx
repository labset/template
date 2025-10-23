import type { PropsWithChildren } from "react";

import { ConnectApiContext } from "./connect-api-context";

export const ConnectApiProvider = ({ children }: PropsWithChildren) => {
  return (
    <ConnectApiContext.Provider value={{}}>
      {children}
    </ConnectApiContext.Provider>
  );
};
