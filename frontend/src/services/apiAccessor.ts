import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:3000",
});

type APIResponse = {
  status: number;
  message: string;
  result: any;
};

const accessor = {
  postRequest: async (path: string, body: any) => {
    return await instance.post<APIResponse>(path, body);
  },
};

export default accessor;
