import { FC, PropsWithChildren } from "react";

// @TODO: "Implement me"

type Props = PropsWithChildren & {
  onDrop: () => void;
};

const Dropzone: FC<Props> = ({ children, onDrop }) => {
  return <div onDrop={() => onDrop()}>{children}</div>;
};

export default Dropzone;
