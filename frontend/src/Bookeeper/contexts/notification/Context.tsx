import { createContext } from "react";
import { NotifierMessage } from "../../components";

type Props = {
  notification: NotifierMessage | null;
  setNotification: (notification: NotifierMessage) => void;
};

const initialStates: Props = {
  notification: null,
  setNotification: () => {},
};

const Context = createContext<Props>(initialStates);

export default Context;
