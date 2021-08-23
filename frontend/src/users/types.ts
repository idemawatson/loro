export interface UserInfo {
  userId: string;
  userName: string;
  numFollowing: number;
  numFollowed: number;
}

export interface getUserInfoRequest {
  userId: string;
}
