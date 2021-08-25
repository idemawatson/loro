import { fetchUserInfoAction } from "users/actions";
import apiAccessor from "services/apiAccessor";
import { getUserInfoRequest, UserInfo } from "users/types";

export const fetchUserInfo = () => {
  return async (dispatch: any) => {
    const request: getUserInfoRequest = {
      userId: "000001",
    };
    const response = await apiAccessor.postRequest("/user", request);
    const userInfo: UserInfo = response.data.result;
    dispatch({
      type: fetchUserInfoAction,
      payload: userInfo,
    });
  };
};
