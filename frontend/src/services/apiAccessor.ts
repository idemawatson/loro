import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:3000",
});

const accessor = {
  postRequest: async (path: string, body: any) => {
    return await instance.post(path, body);
  },
};

export default accessor;
