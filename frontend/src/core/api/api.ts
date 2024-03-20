import axios from "axios";

const instanse = axios.create({
  baseURL: "http://localhost:5000/api",
});

const Api = {
  login(nickname: string, password: string) {
    return instanse.post("/auth/login", {
      username: nickname,
      password,
    });
  },
  register(username: string, password: string, name: string, email: string) {
    return instanse.post("/auth/register", {
      username,
      name,
      email,
      password,
    });
  },
};

export default Api;
