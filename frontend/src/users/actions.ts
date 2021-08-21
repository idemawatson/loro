import actionCreatorFactory from "typescript-fsa";
import { UserInfo } from "users/types";

const actionCreator = actionCreatorFactory();

export const fetchUserInfoAction = actionCreator<UserInfo>("FETCH_USER");
