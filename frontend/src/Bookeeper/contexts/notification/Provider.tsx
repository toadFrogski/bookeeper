import { FC, PropsWithChildren, useState } from "react";
import { NotifierMessage } from "../../components";
import Context from "./Context";

type Props = PropsWithChildren;

const Provider: FC<Props> = ({ children }) => {
  const [notification, setNotification] = useState<NotifierMessage | null>(null);

  return (
    <Context.Provider value={{ notification: notification, setNotification: setNotification }}>
      {children}
    </Context.Provider>
  );
};

export default Provider;
