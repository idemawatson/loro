import { reducerWithInitialState } from "typescript-fsa-reducers";
import * as Actions from "users/actions";
import { UserInfo } from "users/types";

const initialState: UserInfo = {
  userId: "000001",
  userName: "",
  numFollowed: 0,
  numFollowing: 0,
  numPost: 0,
};

const Reducer = reducerWithInitialState(initialState);

export const UserInfoReducer = Reducer.case(
  Actions.fetchUserInfoAction,
  (state, payload: UserInfo) => ({
    ...state,
    ...payload,
  })
);
