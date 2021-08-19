import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:3000",
});

const postRequest = async (path: string, body: any) => {
  return await instance.post(path, body);
};

export default { postRequest };
