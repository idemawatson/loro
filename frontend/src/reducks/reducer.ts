import { reducerWithInitialState } from "typescript-fsa-reducers";
import { UserInfo } from "users/types";

export interface State {
  userInfo: UserInfo;
}

export const initialState: State = {
  userInfo: {
    userId: "000001",
    userName: "",
    numFollowed: 0,
    numFollowing: 0,
  },
};

export const Reducer = reducerWithInitialState(initialState);
