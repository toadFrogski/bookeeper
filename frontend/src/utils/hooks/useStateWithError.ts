import { Dispatch, SetStateAction, useState } from "react";

type UseStateWithErrorReturnType<S> = {
  value: S;
  setValue: Dispatch<SetStateAction<S>>;
  error: string;
  setError: Dispatch<SetStateAction<string>>;
};

function useStateWithError<S>(initialValue: S): UseStateWithErrorReturnType<S> {
  const [value, setValue] = useState<S>(initialValue);
  const [error, setError] = useState("");

  return { value: value, setValue: setValue, error: error, setError: setError };
}

export default useStateWithError;
