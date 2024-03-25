import axios from "axios";

const instanse = axios.create({
  baseURL: "http://localhost:5000/api",
});

const Api = {
  login(nickname: string, password: string) {
    return instanse.post("/auth/login", {
      name: nickname,
      password,
    });
  },
  register(username: string, password: string, email: string) {
    return instanse.post("/auth/register", {
      name: username,
      email,
      password,
    });
  },
};

export default Api;
