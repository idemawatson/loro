import { reducerWithInitialState } from "typescript-fsa-reducers";

export interface State {
  user: {
    userId: string;
  };
}

export const initialState: State = {
  user: {
    userId: "000001",
  },
};

export const Reducer = reducerWithInitialState(initialState);
