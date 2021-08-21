import { fetchUserInfoAction } from "users/actions";
import apiAccessor from "services/apiAccessor";
import { UserInfo } from "users/types";

export const fetchUserInfo = () => {
  return async (dispatch: any) => {
    const response = await apiAccessor.postRequest("/user", "000001");
    const userInfo: UserInfo = response.data;
    dispatch({
      type: fetchUserInfoAction,
      payload: userInfo,
    });
  };
};
